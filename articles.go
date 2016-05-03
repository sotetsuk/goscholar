package main

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
)

const (
	WHOLE_ARTICLE_SELECTOR = ".gs_r"
)

type Articles struct {
	articles []*Article
}

func NewArticles(n int) *Articles {
	as := Articles{}
	as.articles = make([]*Article, n)
	return &as
}

func (as *Articles) ParseAllArticles(doc *goquery.Document, useBibTeX bool) {
	parse := func(i int, s *goquery.Selection) {
		a := NewArticle()
		a.Parse(s, useBibTeX)
		if i >= len(as.articles) { // TODO: fix how to treat slice length
			return
		}
		as.articles[i] = a
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
