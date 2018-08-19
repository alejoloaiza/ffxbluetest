package main

import (
	"fmt"
	"testing"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/valyala/fasthttp"
)

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
