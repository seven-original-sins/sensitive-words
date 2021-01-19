package main

import (
	"github.com/valyala/fasthttp"
	"log"
	"sensitive/config"
	"sensitive/server"
)

func main() {
	cfg := config.GetConfig()

	// new sensitive server
	s, err := server.New(server.BuildAuthorFunc(func(ctx *fasthttp.RequestCtx) bool {
		token := ctx.FormValue("token")
		return string(token) == cfg.Token
	}))
	if err != nil {
		log.Fatal(err)
	}

	s.WatchDictChange()

	// run http server
	if err := fasthttp.ListenAndServe(cfg.ListenAddr, s.HandleFastHTTP); err != nil {
		log.Fatal(err)
	}
}
