package main

import (
	"StaticServer/cmd/routes"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/dgrr/fastws"
	"github.com/valyala/fasthttp"
)

func main() {
	r := fasthttprouter.New()

	fs := &fasthttp.FS{
		Root:       "./public",
		IndexNames: []string{"index.html"},
	}

	r.GET("/", fasthttp.FSHandler(fs.Root, 0))

	r.GET("/ws", fastws.Upgrade(routes.Websocket))

	r.ServeFiles("/public/*filepath", fs.Root)

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
