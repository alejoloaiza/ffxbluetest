package main

type Article struct {
	ID    string   `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
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
type Configuration struct {
	DBUrl       string
	DBViews     string
	DBIndexes   string
	DBViewQuery string
	ServerPort  string
}
