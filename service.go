package main

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

const shortForm = "20060102"
const longForm = "2006-01-02"

func serviceQueryTaggedArticlesByDate(Tag string, Date string) ([]byte, error) {
	if Tag == "" {
		return nil, errors.New("Tag cannot be null on for this endpoint")
	}
	entereddate, err := time.Parse(shortForm, Date)
	if err != nil {
		return nil, errors.New("the given Date is invalid")
	}
	Date = entereddate.Format(longForm)
	result, err := GetTaggedByDate(Date, Tag)
	if err != nil {
		return nil, errors.Wrap(err, "error in the call to the DB method")
	}
	rawresult := RawQueryResponse{}
	err = json.Unmarshal(result, &rawresult)
	if err != nil {
		return nil, errors.Wrap(err, "error while trying to unmarshal response")
	}
	if len(rawresult.Rows) == 0 {
		return nil, errors.New("no records found for the given parameters")
	}
	finalresult := transformRawResponseToFinal(rawresult)
	return finalresult, nil

}
func serviceQueryArticleByID(ID string) ([]byte, error) {
	if ID == "" {
		return nil, errors.New("ID cannot be null on article search")
	}
	resp, err := GetDocumentByID(ID)
	if err != nil {
		return nil, errors.Wrap(err, "error in the call to the DB method")
	}
	draftArticle := Article{}
	json.Unmarshal(resp, &draftArticle)
	if draftArticle.validate() != nil {
		return nil, errors.New("document not found")
	}
	return draftArticle.toBytes(), nil
}
func transformRawResponseToFinal(raw interface{}) []byte {
	oldresp, ok := raw.(RawQueryResponse)
	if !ok {
		return []byte("Error converting query response\n")
	}
	newresp := FinalResponse{}
	newresp.Count = len(oldresp.Rows)
	if len(oldresp.Rows) > 0 {
		newresp.Tag = oldresp.Rows[0].Key[1]
	}
	ids := []string{}
	related := make(map[string]bool)
	for _, row := range oldresp.Rows {
		ids = append(ids, row.ID)
		for _, tags := range row.Value {
			related[tags] = true
		}
	}
	if len(ids) > Localconfig.MaxArticlesInQuery {
		ids = ids[len(ids)-Localconfig.MaxArticlesInQuery:]
	}
	relatedtags := []string{}
	for tag := range related {
		if tag != newresp.Tag {
			relatedtags = append(relatedtags, tag)
		}
	}
	newresp.Articles = ids
	newresp.RelatedTags = relatedtags
	bytesnewresp, _ := json.Marshal(newresp)
	return bytesnewresp
}

func serviceCreateArticle(body []byte) ([]byte, error) {
	draftArticle := Article{}
	err := json.Unmarshal(body, &draftArticle)
	if err != nil {
		return nil, errors.Wrap(err, "error while parsing the json, invalid format")
	}
	if err = draftArticle.validate(); err != nil {
		return nil, errors.Wrap(err, "validation failed, mandatory fields missing")
	}
	dbresult, err := InsertArticle(draftArticle.toString(), draftArticle.ID)
	if err != nil {
		return nil, errors.Wrap(err, "error encountered while inserting into DB")
	}
	return dbresult, nil
}

func (a *Article) toBytes() []byte {
	bytes, _ := json.Marshal(a)
	return bytes
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
