package main

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/valyala/fasthttp"
)

var DBSetup bool = false

func Save(doc string, id string) ([]byte, error) {
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
		return nil, errors.Wrap(err, "error while inserting in DB")
	}
	bodyBytes := resp.Body()
	dbresponse := CouchDBResponse{}
	json.Unmarshal(bodyBytes, &dbresponse)
	if !dbresponse.Ok {
		return nil, errors.New("error while inserting the new record on the DB")
	}
	return bodyBytes, nil

}

func SetupDB() bool {
	db, _ := CreateDB()
	index, _ := CreateIndex()
	views, _ := CreateViews()
	if db && index && views {
		return true
	}
	return false
}

func CreateDB() (bool, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(Localconfig.DBUrl)
	req.Header.SetMethodBytes([]byte("PUT"))
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		return false, errors.Wrap(err, "error while creating DB")
	}
	dbresponse := CouchDBResponse{}
	json.Unmarshal(resp.Body(), &dbresponse)
	if !dbresponse.Ok {
		return false, errors.New("error while creating DB")
	}
	return true, nil
}

func CreateIndex() (bool, error) {
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
		return false, errors.Wrap(err, "error while creating index in DB")
	}
	dbresponse := CouchDBResponse{}
	json.Unmarshal(resp.Body(), &dbresponse)
	if !dbresponse.Ok {
		return false, errors.New("error while creating index in DB")
	}
	return true, nil
}
func CreateViews() (bool, error) {
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
		return false, errors.Wrap(err, "error while creating views in DB")
	}
	dbresponse := CouchDBResponse{}
	json.Unmarshal(resp.Body(), &dbresponse)
	if !dbresponse.Ok {
		return false, errors.New("error while creating views in DB")
	}
	return true, nil
}

func GetDocumentByID(id string) ([]byte, error) {
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
		return nil, errors.Wrap(err, "error while getting article from DB")
	}
	bodyBytes := resp.Body()
	return bodyBytes, nil
}

func GetTaggedByDate(date string, tag string) ([]byte, error) {
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
		return nil, errors.Wrap(err, "error while getting tagged by date from DB")
	}
	bodyBytes := resp.Body()
	return bodyBytes, nil
}
