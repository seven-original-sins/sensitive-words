package main

import (
	"github.com/valyala/fasthttp"
	"log"
	"sensitive/config"
)

func main() {
	cfg := config.GetConfig()
	// init server
	s := new(Server)
	// register author
	author := BuildAuthorFunc(func(ctx *fasthttp.RequestCtx) bool {
		token := ctx.FormValue("token")
		return string(token) == cfg.Token
	})
	err := s.Build(author)
	defer s.Close()
	if err != nil {
		log.Fatal(err)
	}

	s.WatchDictChange()

	// run http server
	if err := fasthttp.ListenAndServe(cfg.ListenAddr, s.HandleFastHTTP); err != nil {
		log.Fatal(err)
	}
}
