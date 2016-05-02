package go_scholar

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
)

const (
	WHOLE_ARTICLE_SELECTOR = ".gs_r"
)

type Articles struct {
	articles []Article
}

func NewArticles(n int) *Articles {
	as := Articles{}
	as.articles = make([]Article, n)
	return &as
}

func (as *Articles) ParseAllArticles(doc *goquery.Document, useBibTeX bool) {
	parse := func(i int, s *goquery.Selection) {
		as.articles[i].Parse(s, useBibTeX)
	}
	doc.Find(WHOLE_ARTICLE_SELECTOR).Each(parse)
}

func (as *Articles) Json() string {
	bytes, err := json.Marshal(as.articles)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}
