package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

const (
	WHOLE_ARTICLE_SELECTOR = ".gs_r"
)

type Articles struct {
	articles chan *Article
}

func NewArticles(buffer_size int) *Articles {
	as := Articles{}
	as.articles = make(chan *Article, buffer_size)
	return &as
}

func (as *Articles) ParseAllArticles(doc *goquery.Document, useBibTeX bool) {
	defer close(as.articles)

	parse := func(i int, s *goquery.Selection) {
		a := NewArticle()
		a.Parse(s, useBibTeX)

		// Add this Article to Articles
		if a.IsValid() {
			as.articles <- a
		}
	}
	doc.Find(WHOLE_ARTICLE_SELECTOR).Each(parse)
}

func (as *Articles) StdoutJson() {
	fmt.Printf("[")
	initial_flag := true
	for a := range as.articles {
		if initial_flag {
			fmt.Print(a.Json())
			initial_flag = false
		} else {
			fmt.Print(",", a.Json())
		}
	}
	fmt.Println("]")
}