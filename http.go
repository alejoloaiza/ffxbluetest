package main

import (
	"fmt"

	"github.com/buaazp/fasthttprouter"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

func main() {
	loadConfig("./config/config.json")
	createRestServer()
}

func createRestServer() error {
	router := fasthttprouter.New()
	router.POST("/articles", createArticle)
	router.GET("/articles/:id", queryArticleByID)
	router.GET("/tags/:tags/:date", queryTaggedByDate)
	fmt.Printf("Server started! on port %s\n", Localconfig.ServerPort)
	err := fasthttp.ListenAndServe(Localconfig.ServerPort, router.Handler)
	if err != nil {
		return errors.New("error while trying to listen on the given port")
	}
	return nil
}
func createArticle(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	response, err := serviceCreateArticle(body)
	if err != nil {
		ctx.Response.SetStatusCode(500)
		fmt.Fprint(ctx, err)
	}
	ctx.Response.SetStatusCode(200)
	fmt.Fprint(ctx, string(response))
}
func queryArticleByID(ctx *fasthttp.RequestCtx) {
	ID, ok := ctx.UserValue("id").(string)
	if !ok {
		ctx.Response.SetStatusCode(500)
		fmt.Fprint(ctx, "Error while fetching the input id\n")
	}
	resp, err := serviceQueryArticleByID(ID)
	if err != nil {
		ctx.Response.SetStatusCode(500)
		fmt.Fprint(ctx, err)
	}
	ctx.Response.SetStatusCode(200)
	fmt.Fprint(ctx, string(resp))
}
func queryTaggedByDate(ctx *fasthttp.RequestCtx) {
	Tag, ok := ctx.UserValue("tags").(string)
	if !ok {
		ctx.Response.SetStatusCode(500)
		fmt.Fprint(ctx, "Error while fetching the input id\n")
	}
	Date, ok := ctx.UserValue("date").(string)
	if !ok {
		ctx.Response.SetStatusCode(500)
		fmt.Fprint(ctx, "Error while fetching the input id\n")
	}
	result, err := serviceQueryTaggedArticlesByDate(Tag, Date)
	if err != nil {
		ctx.Response.SetStatusCode(500)
		fmt.Fprint(ctx, err)
	}
	ctx.Response.SetStatusCode(200)
	fmt.Fprint(ctx, string(result))
}
