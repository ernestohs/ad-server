package main

import (
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"
)

func main() {

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/ad/search":
			fmt.Println("search")
		case "/ad/impression":
			fmt.Println("impression")
		case "/ad/click":
			fmt.Println("click")
		case "/ad/conversion":
			fmt.Println("conversion")
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}

	port := ":" + strconv.Itoa(8001)
	fasthttp.ListenAndServe(port, requestHandler)
}
