package goscholar

import (
	"testing"
)

func TestSearchUrl(t *testing.T) {
	expected := "https://scholar.google.co.jp/scholar?hl=en&q=deep+learning+author:\"y+bengio\"&as_ylo=2015&as_yhi=&num=100&start=20"
	q := &Query{Author: "y bengio", Title: "", Keywords: "deep learning", After: "2015", Before: "", Num: "100", Start: "20"}

	if actual := q.SearchUrl(); actual != expected {
		t.Error(testErr{expected: expected, actual: actual})
	}
}

func TestFindUrl(t *testing.T) {
	expected := "https://scholar.google.co.jp/scholar?hl=en&cluster=5362332738201102290&num=1"
	q := &Query{ClusterId: "5362332738201102290"}

	if actual := q.FindUrl(); actual != expected {
		t.Error(testErr{expected: expected, actual: actual})
	}
}

func TestCiteUrl(t *testing.T) {
	expected := "https://scholar.google.co.jp/scholar?hl=en&cites=5362332738201102290&as_ylo=2012&as_yhi=&num=40&start=20"
	q := &Query{ClusterId: "5362332738201102290", After: "2012", Num: "40", Start: "20"}

	if actual := q.CiteUrl(); actual != expected {
		t.Error(testErr{expected: expected, actual: actual})
	}
}

func TestCitePopUpQueryUrl(t *testing.T) {
	expected := "https://scholar.google.co.jp/scholar?q=info:XOJff8gPiHAJ:scholar.google.com/&output=cite&scirp=0&hl=en"
	q := &Query{InfoId: "XOJff8gPiHAJ"}

	if actual := q.CitePopUpQueryUrl(); actual != expected {
		t.Error(testErr{expected: expected, actual: actual})
	}
}
