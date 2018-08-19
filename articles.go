package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

const shortForm = "20060102"
const longForm = "2006-01-02"

func queryTaggedArticlesByDate(ctx *fasthttp.RequestCtx) {
	Tag, ok := ctx.UserValue("tags").(string)
	if !ok {
		fmt.Fprint(ctx, "Error converting Tag to string\n")
	}
	if Tag == "" {
		fmt.Fprint(ctx, "Tag is empty, please give one Tag\n")
	}
	Date, ok := ctx.UserValue("date").(string)
	if !ok {
		fmt.Fprint(ctx, "Error converting Date to string\n")
	}
	entereddate, err := time.Parse(shortForm, Date)
	if err != nil {
		fmt.Fprint(ctx, "Date is invalid\n")
	}
	Date = entereddate.Format(longForm)
	result := GetTaggedByDate(Date, Tag)
	rawresult := RawQueryResponse{}
	err = json.Unmarshal(result, &rawresult)
	if err != nil {
		fmt.Fprint(ctx, "Error parsing response")
	}
	finalresult := transformRawResponseToFinal(rawresult)
	fmt.Fprint(ctx, string(finalresult))

}

func transformRawResponseToFinal(raw interface{}) []byte {
	oldresp, ok := raw.(RawQueryResponse)
	if !ok {
		return []byte("Error converting query response\n")
	}
	newresp := FinalResponse{}
	newresp.Count = oldresp.Offset
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

func queryArticleByID(ctx *fasthttp.RequestCtx) {
	ID, ok := ctx.UserValue("id").(string)
	if !ok {
		fmt.Fprint(ctx, "Error converting id to string\n")
	}
	if ID == "" {
		fmt.Fprint(ctx, "ID is empty, please give one id\n")
	}
	resp := GetDocumentByID(ID)
	draftArticle := Article{}
	json.Unmarshal(resp, &draftArticle)
	fmt.Fprint(ctx, draftArticle.toString())
}

func createArticle(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	draftArticle := Article{}
	err := json.Unmarshal(body, &draftArticle)
	if err != nil {
		fmt.Fprint(ctx, "Error while parsing the json, invalid format\n")
	}
	if err = draftArticle.validate(); err != nil {
		fmt.Fprintf(ctx, "Error %s\n", err)
	}
	Save(draftArticle.toString(), draftArticle.ID)
	fmt.Fprint(ctx, "Ok\n")
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
