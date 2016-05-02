package go_scholar

import (
	"testing"
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

func TestParseTitle(t *testing.T) {
	// fetch goquery.Document
	url := "https://scholar.google.co.jp/scholar?hl=en&q=\"Learning+deep+architectures+for+AI\"&as_ylo=&as_yhi=&start=&num="
	doc, err := goquery.NewDocument(url)
	if err != nil {
		t.Error(fmt.Sprintf("URL is not valid: %v", err.Error()))
	}

	a := NewArticle()
	a.parseTitle(doc.Find(WHOLE_ARTICLE_SELECTOR).First())
	expected := "Learning deep architectures for AI"
	if a.Title != expected {
		t.Error(fmt.Sprintf("Expected: %v\nActual: %v", a.Title, expected))
	}
	fmt.Println("Title: ", a.Title)
}
