package main

import (
	"fmt"
	"testing"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/valyala/fasthttp"
)

var ArticleID string

const Endpoint = "http://localhost:3000/articles"

func BenchmarkArticles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u, _ := uuid.NewV4()
		Input := `{
			"id": "%s",
			"title": "My Article about Health",
			"date" : "2016-09-22",
			"body" : "some text, potentially containing simple markup about how potato chips are great",
			"tags" : ["health", "fitness", "science"]
		  }`
		request := fmt.Sprintf(Input, u)
		req := fasthttp.AcquireRequest()
		req.SetRequestURI(Endpoint)
		req.Header.SetContentType("application/json")
		req.Header.SetMethodBytes([]byte("POST"))
		req.SetBodyString(request)
		resp := fasthttp.AcquireResponse()
		client := &fasthttp.Client{}
		if err := client.Do(req, resp); err != nil {
			fmt.Println("Error:", err.Error())
			b.Error(err)
		}

	}

}
func TestPositiveCaseCreateArticle(t *testing.T) {
	loadConfig("./config/config.json")
	var ctx fasthttp.RequestCtx
	u, _ := uuid.NewV4()
	Input := `{
		"id": "%s",
		"title": "My Article about Health",
		"date" : "2016-09-22",
		"body" : "some text, potentially containing simple markup about how potato chips are great",
		"tags" : ["health", "fitness", "science"]
	  }`
	request := fmt.Sprintf(Input, u)
	ArticleID = u.String()
	ctx.Request.SetBodyString(request)
	createArticle(&ctx)
	if ctx.Response.StatusCode() != 200 {
		t.Fail()
	}
}
func TestNegativeCaseCreateArticle(t *testing.T) {
	loadConfig("./config/config.json")
	var ctx fasthttp.RequestCtx
	Input := `{
		"id": "%s",
		"title": "My Article about Health",
		"date" : "2016-09-22",
		"body" : "some text, potentially containing simple markup about how potato chips are great",
		"tags" : ["health", "fitness", "science"]
	  }`
	request := fmt.Sprintf(Input, ArticleID)
	ctx.Request.SetBodyString(request)
	createArticle(&ctx)
	if ctx.Response.StatusCode() == 200 {
		t.Fail()
	}
}

func TestPositiveCaseGetArticle(t *testing.T) {
	loadConfig("./config/config.json")
	var ctx fasthttp.RequestCtx
	ctx.SetUserValue("id", ArticleID)
	queryArticleByID(&ctx)
	if ctx.Response.StatusCode() != 200 {
		t.Fail()
	}
}
func TestPositiveCaseTaggedArticle(t *testing.T) {
	loadConfig("./config/config.json")
	var ctx fasthttp.RequestCtx
	ctx.SetUserValue("tags", "fitness")
	ctx.SetUserValue("date", "20160922")
	queryTaggedByDate(&ctx)
	if ctx.Response.StatusCode() != 200 {
		t.Fail()
	}
}
func TestNegativeCase1TaggedArticle(t *testing.T) {
	loadConfig("./config/config.json")
	var ctx fasthttp.RequestCtx
	ctx.SetUserValue("tag", "fitness")
	ctx.SetUserValue("date", "2016-13-22")
	queryTaggedByDate(&ctx)
	if ctx.Response.StatusCode() == 200 {
		t.Fail()
	}
}
func TestNegativeCase2TaggedArticle(t *testing.T) {
	loadConfig("./config/config.json")
	var ctx fasthttp.RequestCtx
	ctx.SetUserValue("tags", "fitness")
	ctx.SetUserValue("date", "2016-13-22")
	queryTaggedByDate(&ctx)
	if ctx.Response.StatusCode() == 200 {
		t.Fail()
	}
}
func TestNegativeCase1GetArticle(t *testing.T) {
	loadConfig("./config/config.json")
	var ctx fasthttp.RequestCtx
	ctx.SetUserValue("id", "xxx")
	queryTaggedByDate(&ctx)
	if ctx.Response.StatusCode() == 200 {
		t.Fail()
	}
}
func TestNegativeCaseConfig(t *testing.T) {
	loadConfig("./config/config2.json")
	if Localconfig != nil {
		t.Fail()
	}
}
func TestNegativeServer(t *testing.T) {
	err := createRestServer()
	if err == nil {
		t.Fail()
	}
}
