package goscholar

import (
	"testing"
)

func TestSearchUrl(t *testing.T) {
	expected := "https://scholar.google.co.jp/scholar?hl=en&q=deep+learning+author:\"y+bengio\"&as_ylo=2015&as_yhi=&num=100&start=20"
	q := &Query{author: "y bengio", title: "", keywords:"deep learning", after:"2015", before:"", num:"100", start:"20"}

	if actual := q.SearchUrl(); actual != expected {
		t.Error(TestErr{expected:expected, actual:actual})
	}
}

func TestFindUrl(t *testing.T) {
	expected := "https://scholar.google.co.jp/scholar?hl=en&cluster=8108748482885444188&num=1"
	q := &Query{cluster_id:"8108748482885444188"}

	if actual := q.FindUrl(); actual != expected {
		t.Error(TestErr{expected:expected, actual:actual})
	}
}

func TestCiteUrl(t *testing.T) {
	expected := "https://scholar.google.co.jp/scholar?hl=en&cites=8108748482885444188&as_ylo=2012&as_yhi=&num=40&start=20"
	q := &Query{cluster_id:"8108748482885444188", after:"2012", num:"40", start:"20"}

	if actual := q.CiteUrl(); actual != expected {
		t.Error(TestErr{expected:expected, actual:actual})
	}
}

func TestCitePopUpQueryUrl(t *testing.T) {
	expected := "https://scholar.google.co.jp/scholar?q=info:XOJff8gPiHAJ:scholar.google.com/&output=cite&scirp=0&hl=en"
	q := &Query{info_id:"XOJff8gPiHAJ"}

	if actual := q.CitePopUpQueryUrl(); actual != expected {
		t.Error(TestErr{expected:expected, actual:actual})
	}
}
