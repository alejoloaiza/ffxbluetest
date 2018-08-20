package main

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/valyala/fasthttp"
)

// DBSetup is used to control whether setup is required or not.
var DBSetup = false

// InsertArticle is used to store an article inside the DB.
func InsertArticle(doc string, id string) ([]byte, error) {
	if !DBSetup && SetupDB() {
		DBSetup = true
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

// SetupDB will create the DB, indexes and views.
func SetupDB() bool {
	db, _ := CreateDB()
	index, _ := CreateIndex()
	views, _ := CreateViews()
	if db && index && views {
		return true
	}
	return false
}

// CreateDB connects to couchdb and creates the DB where all documents will be stored.
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

// CreateIndex will create an index on the tag array field and date field.
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

// CreateViews creates a view inside the couchdb to extract the data by date and tags.
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

// GetDocumentByID is used to extract from DB a particular Article.
func GetDocumentByID(id string) ([]byte, error) {
	if !DBSetup && SetupDB() {
		DBSetup = true
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

// GetTaggedByDate is used to extract from DB the information of tagged articles on a given date.
func GetTaggedByDate(date string, tag string) ([]byte, error) {
	if !DBSetup && SetupDB() {
		DBSetup = true
	}
	req := fasthttp.AcquireRequest()
	RequestURL := fmt.Sprintf(Localconfig.DBUrl+Localconfig.DBViewQuery, date, tag, date, tag)
	req.SetRequestURI(RequestURL)
	req.Header.SetMethodBytes([]byte("GET"))
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		return nil, errors.Wrap(err, "error while getting tagged by date from DB")
	}
	bodyBytes := resp.Body()
	return bodyBytes, nil
}
