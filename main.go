package main

import (
	"github.com/valyala/fasthttp"
	"log"
	"sensitive/config"
)

func main() {
	// init server
	s := new(Server)
	err := s.Build()
	defer s.Close()
	if err != nil {
		log.Fatal(err)
	}

	s.WatchDictChange()

	cfg := config.GetConfig()

	// run http server
	if err := fasthttp.ListenAndServe(cfg.ListenAddr, s.HandleFastHTTP); err != nil {
		log.Fatal(err)
	}
}
