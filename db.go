package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

var setup bool = false

const dbUrl = "http://couchdb:5984/articles"
const dbView = "/_design/myviews"
const dbViewQuery = "/_design/myviews/_view/tags?startkey=['%s','%s']&endkey=startkey=['%s','%s']"

func Save(doc string, id string) {
	if !setup {
		if SetupDB() {
			setup = true
		}
	}
	req := fasthttp.AcquireRequest()
	RequestURL := dbUrl + "/" + id
	req.SetRequestURI(RequestURL)
	req.Header.SetMethodBytes([]byte("PUT"))
	req.SetBodyString(doc)
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		println("Error:", err.Error())
	} else {
		bodyBytes := resp.Body()
		println(string(bodyBytes))
	}
}
func SetupDB() bool {
	if CreateDB() && CreateIndex() && CreateViews() {
		return true
	}
	return false
}
func CreateDB() bool {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(dbUrl)
	req.Header.SetMethodBytes([]byte("PUT"))
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		println("Error:", err.Error())
	} else {
		bodyBytes := resp.Body()
		println(string(bodyBytes))
		return true
	}
	return false
}

func CreateIndex() bool {
	return true
}
func CreateViews() bool {
	view := `{
		"language": "javascript",
		"views": {
			"tags": {
				"map": "function(doc) {    var tag, i;    if (doc.tags && doc.id) {        for (i in doc.tags) {            tag = doc.tags[i];            emit([doc.date, tag], doc.tags);        }    }}"
			}
		}
	}`
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(dbUrl + dbView)
	req.Header.SetMethodBytes([]byte("PUT"))
	req.SetBodyString(view)
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		println("Error:", err.Error())
	} else {
		bodyBytes := resp.Body()
		println(string(bodyBytes))
		return true
	}
	return false
}

func GetDocumentByID(id string) []byte {
	req := fasthttp.AcquireRequest()
	RequestURL := fmt.Sprintf("%s/%s", dbUrl, id)
	req.SetRequestURI(RequestURL)
	req.Header.SetMethodBytes([]byte("GET"))

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		println("Error:", err.Error())
	} else {
		bodyBytes := resp.Body()
		return bodyBytes
	}
	return []byte("")
}

func GetTaggedByDate(date string, tag string) []byte {
	req := fasthttp.AcquireRequest()

	RequestURL := fmt.Sprintf(dbUrl+dbViewQuery, date, tag, date, tag)
	req.SetRequestURI(RequestURL)
	req.Header.SetMethodBytes([]byte("GET"))

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		println("Error:", err.Error())
	} else {
		bodyBytes := resp.Body()
		return bodyBytes
	}
	return []byte("")
}
