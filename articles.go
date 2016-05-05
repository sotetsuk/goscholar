package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

const (
	WHOLE_ARTICLE_SELECTOR = ".gs_r"
)

func ParseArticles(ch chan *Article, doc *goquery.Document) {
	defer close(ch)

	parse := func(i int, s *goquery.Selection) {
		a := NewArticle()
		a.Parse(s)

		// Add this Article to Articles
		if a.isValid() {
			ch <- a
		}
	}
	doc.Find(WHOLE_ARTICLE_SELECTOR).Each(parse)
}

func StdoutArticleAsJson(ch chan *Article) {
	fmt.Printf("[")
	initial_flag := true
	for a := range ch {
		if initial_flag {
			fmt.Print(a.Json())
			initial_flag = false
		} else {
			fmt.Print(",", a.Json())
		}
	}
	fmt.Println("]")
}