package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var (
	addr = flag.String("addr", ":3000", "TCP address to listen to")
)

func main() {
	flag.Parse()
	createRestServer()

}

func queryTaggedArticlesByDate(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Nothing for now, just dummy\n")
}

func queryArticleByID(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "dummy for now, %s!\n", ctx.UserValue("id"))
}
func createArticle(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Nothing for now, just dummy\n")
}

func createRestServer() {
	router := fasthttprouter.New()
	router.POST("/articles", createArticle)
	router.GET("/articles/:id", queryArticleByID)
	router.GET("/tags/{tagName}/{date}", queryTaggedArticlesByDate)

	log.Fatal(fasthttp.ListenAndServe(*addr, router.Handler))
}
