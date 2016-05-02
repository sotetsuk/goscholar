package main

import (
	"testing"
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

func TestArticles(t *testing.T) {
	url := "https://scholar.google.co.jp/scholar?hl=en&q=deep+learning+author:\"Bengio\"&as_ylo=2015&as_yhi=&start=20&num=100"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		t.Error(fmt.Sprintf("URL is not valid: %v", err.Error()))
	}

	as := NewArticles(10)
	as.ParseAllArticles(doc, false)
	for _, a := range as.articles {
		fmt.Println("============================================")
		if a != nil {
			a.dump()
		}
	}

	// TODO: write test
}

