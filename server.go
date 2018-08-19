package main

import (
	"encoding/json"
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
	ID, ok := ctx.UserValue("id").(string)
	if !ok {
		fmt.Fprint(ctx, "Error converting id to string\n")
	}
	if ID == "" {
		fmt.Fprint(ctx, "ID is empty, please give one id\n")
	}
	resp := GetDocumentByID(ID)
	draftArticle := Article{}
	json.Unmarshal(resp, &draftArticle)
	fmt.Fprint(ctx, draftArticle.toString())
}

func createArticle(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	draftArticle := Article{}
	err := json.Unmarshal(body, &draftArticle)
	if err != nil {
		fmt.Fprint(ctx, "Error while parsing the json, invalid format\n")
	}
	if err = draftArticle.validate(); err != nil {
		fmt.Fprintf(ctx, "Error %s\n", err)
	}
	Save(draftArticle.toString(), draftArticle.ID)
	fmt.Fprint(ctx, "Ok\n")
}

func createRestServer() {
	router := fasthttprouter.New()
	router.POST("/articles", createArticle)
	router.GET("/articles/:id", queryArticleByID)
	router.GET("/tags/:tags/:date", queryTaggedArticlesByDate)
	fmt.Println("Server started!")

	log.Fatal(fasthttp.ListenAndServe(*addr, router.Handler))
}
