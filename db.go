package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

var DBSetup bool = false

func Save(doc string, id string) {
	if !DBSetup {
		if SetupDB() {
			DBSetup = true
		}
	}
	req := fasthttp.AcquireRequest()
	RequestURL := Localconfig.DBUrl + "/" + id
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
	req.SetRequestURI(Localconfig.DBUrl)
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
	index := `{
		"index": {
		  "fields": [
			"_id",
			"tags.[]",
		  "date"
		  ]
		},
		"type": "json"
	  }`
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(Localconfig.DBUrl + Localconfig.DBIndexes)
	req.Header.SetContentType("application/json")
	req.Header.SetMethodBytes([]byte("POST"))
	req.SetBodyString(index)
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
	req.SetRequestURI(Localconfig.DBUrl + Localconfig.DBViews)
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
	if !DBSetup {
		if SetupDB() {
			DBSetup = true
		}
	}
	req := fasthttp.AcquireRequest()
	RequestURL := fmt.Sprintf("%s/%s", Localconfig.DBUrl, id)
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
	if !DBSetup {
		if SetupDB() {
			DBSetup = true
		}
	}
	req := fasthttp.AcquireRequest()

	RequestURL := fmt.Sprintf(Localconfig.DBUrl+Localconfig.DBViewQuery, date, tag, date, tag)
	req.SetRequestURI(RequestURL)
	fmt.Println(RequestURL)
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
