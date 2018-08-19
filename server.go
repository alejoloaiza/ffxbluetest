package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	GetConfig("./config/config.json")
	createRestServer()

}

func createRestServer() {
	router := fasthttprouter.New()
	router.POST("/articles", createArticle)
	router.GET("/articles/:id", queryArticleByID)
	router.GET("/tags/:tags/:date", queryTaggedArticlesByDate)
	fmt.Printf("Server started! on port %s\n", Localconfig.ServerPort)
	log.Fatal(fasthttp.ListenAndServe(Localconfig.ServerPort, router.Handler))
}
