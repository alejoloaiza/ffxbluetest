package main

// Article is the structs that defines an article for the API.
type Article struct {
	ID    string   `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

// RawQueryResponse is used to capture the Couchdb response for the view query.
type RawQueryResponse struct {
	TotalRows int `json:"total_rows"`
	Offset    int `json:"offset"`
	Rows      []struct {
		ID    string   `json:"id"`
		Key   []string `json:"key"`
		Value []string `json:"value"`
	} `json:"rows"`
}

// CouchDBResponse is mainly used to check the response given by CouchDB.
type CouchDBResponse struct {
	Ok bool `json:"ok"`
}

// FinalResponse is the final struct to deliver the result of the query by tag and date.
type FinalResponse struct {
	Tag         string   `json:"tag"`
	Count       int      `json:"count"`
	Articles    []string `json:"articles"`
	RelatedTags []string `json:"related_tags"`
}

// Configuration is used to store all the configuration fields from the config.json
type Configuration struct {
	DBUrl              string
	DBViews            string
	DBIndexes          string
	DBViewQuery        string
	ServerPort         string
	MaxArticlesInQuery int
}
