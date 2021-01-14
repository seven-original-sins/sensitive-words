package tools

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)

func WriteJSON(ctx *fasthttp.RequestCtx, v interface{}) {
	j, err := json.Marshal(v)
	if err != nil {
		ctx.Error("server error", 500)
		log.Printf("json encode error: %s", err)
		return
	}

	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))

	ctx.SetBody(j)
}
