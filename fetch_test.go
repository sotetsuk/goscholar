package goscholar

import (
	"testing"
)

func TestFetch(t *testing.T) {
	url := "https://scholar.google.co.jp/scholar?hl=en&cluster=5362332738201102290&num=1"

	doc, err := Fetch(url)
	if err != nil {
		t.Skip(err)
	}

	expected := "Deep learning"
	actual := doc.Find(WHOLE_ARTICLE_SELECTOR).First().Find(ARTICLE_TITLE_SELECTOR).Text()

	if actual != expected {
		t.Error(TestErr{expected:expected, actual:actual})
	}
}
