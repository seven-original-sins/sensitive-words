package server

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/anknown/ahocorasick"
	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"sensitive/config"
	"sensitive/tools"
	"sync"
	"time"
)

type Server struct {
	count  int
	db     *sql.DB
	m      *goahocorasick.Machine
	rw     sync.RWMutex
	author Author
}

// request author
type Author interface {
	Auth(ctx *fasthttp.RequestCtx) bool
}

// build DB and Machine.
func (s *Server) Build(author Author) error {
	var err error
	cfg := config.GetConfig()
	s.author = author

	s.db, err = sql.Open("mysql", cfg.DBUsername + ":" + cfg.DBPassword + "@(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBDatabase)
	if err != nil {
		return fmt.Errorf("cannot connect to mysql: %s", err)
	}

	err = s.LoadWords()
	if err != nil {
		return fmt.Errorf("laod words error: %s", err)
	}

	return nil
}

func (s *Server) Close() {
	if err :=s.db.Close(); err != nil {
		log.Printf("close db: %s", err)
	}
}

// watch dict change.
// pre 5 seconds
func (s *Server) WatchDictChange() {
	go func() {
		for {
			count := 0
			row := s.db.QueryRow("select count(*) from t_sensitive_word where _iDeleteTime = 0")

			if err := row.Scan(&count); err != nil {
				time.Sleep(time.Second * 5)
				continue
			}

			s.rw.RLock()
			sc := s.count
			s.rw.RUnlock()
			if count != sc {
				_ = s.LoadWords()
			}


			time.Sleep(time.Second * 5)
		}
	}()
}

// load words from database.
func (s *Server) LoadWords() error {
	rows, err := s.db.Query("select sWord from t_sensitive_word where _iDeleteTime = 0")
	if err != nil {
		return fmt.Errorf("get words error: %s", err)
	}
	defer rows.Close()

	dict := make([][]rune, 0, 1024)
	for rows.Next() {
		var word string
		if err := rows.Scan(&word); err != nil {
			log.Println(err)
			continue
		}

		dict = append(dict, []rune(word))
	}

	if len(dict) == 0 {
		return fmt.Errorf("dict cannot be null")
	}

	s.rw.Lock()
	defer s.rw.Unlock()
	s.count = len(dict)
	s.m = new(goahocorasick.Machine)
	if err := s.m.Build(dict); err != nil {
		return err
	}

	return nil
}

// handle http request.
func (s *Server) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	// authorization
	if s.author != nil {
		result := s.author.Auth(ctx)
		if !result {
			ctx.Error("unauthorized.", http.StatusForbidden)
			return
		}
	}
	switch string(ctx.Path()) {
	case "/words/reload":
		s.reloadWords(ctx)
	case "/words/search":
		s.search(ctx)
	}
}

type ReloadWordsResponse struct {
	Result bool `json:"result"`
	Msg string `json:"msg"`
}

// reload dict.
func (s *Server) reloadWords(ctx *fasthttp.RequestCtx) {
	err := s.LoadWords()
	result := true
	msg := ""
	if err != nil {
		result = false
		msg = err.Error()
	}

	tools.WriteJSON(ctx, &ReloadWordsResponse{Result: result, Msg: msg})
}

// /words/search request handler.
func (s *Server) search(ctx *fasthttp.RequestCtx) {
	response := new(SearchResponse)
	result := make([]*HitWord, 0)
	response.Result = result
	content := bytes.Runes(ctx.FormValue("content"))
	if len(content) == 0 {
		response.Type = "PASS"
		tools.WriteJSON(ctx, response)
		return
	}

	s.rw.RLock()
	terms := s.m.MultiPatternSearch(content, false)
	s.rw.RUnlock()

	for _, term := range terms {
		result = append(result,&HitWord{Pos: term.Pos, Word: string(term.Word)})
	}

	if len(result) == 0 {
		response.Type = "PASS"
	} else {
		response.Type = "HIT"
		response.Result = result
	}

	tools.WriteJSON(ctx, response)
}

type SearchResponse struct {
	Type string       `json:"type"`
	Result []*HitWord `json:"result"`
}

type HitWord struct {
	Pos int `json:"pos"`
	Word string `json:"word"`
}

type AuthorFunc func(ctx *fasthttp.RequestCtx) bool

func (f AuthorFunc) Auth(ctx *fasthttp.RequestCtx) bool {
	return f(ctx)
}

func BuildAuthorFunc(f func(*fasthttp.RequestCtx) bool) AuthorFunc {
	return f
}

func New(author Author) (*Server, error) {
	// init server
	s := new(Server)
	err := s.Build(author)
	defer s.Close()
	if err != nil {
		return nil, err
	}

	return s, nil
}