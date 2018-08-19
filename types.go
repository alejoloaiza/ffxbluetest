package main

import (
	"encoding/json"
	"errors"
)

type Article struct {
	ID    string   `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

func (a *Article) toString() string {
	str, _ := json.Marshal(a)
	return string(str)
}

func (a *Article) validate() error {
	if a.ID == "" || a.Title == "" || a.Date == "" || a.Body == "" || len(a.Tags) == 0 {
		return errors.New("Cannot insert document, Mandatory field validation")
	}
	return nil
}

type CouchResponse struct {
	Ok  bool   `json:"ok"`
	ID  string `json:"id"`
	Rev string `json:"rev"`
}

type RawQueryResponse struct {
	TotalRows int `json:"total_rows"`
	Offset    int `json:"offset"`
	Rows      []struct {
		ID    string   `json:"id"`
		Key   []string `json:"key"`
		Value []string `json:"value"`
	} `json:"rows"`
}
type FinalResponse struct {
	Tag         string   `json:"tag"`
	Count       int      `json:"count"`
	Articles    []string `json:"articles"`
	RelatedTags []string `json:"related_tags"`
}
