package go_scholar

import (
	"testing"
	"fmt"
)

func TestSearchQuery(t *testing.T) {
	// exec SearchQuery()
	query, err := SearchQuery("deep learning", "Bengio", "", "2015", "", "20", "100")
	if err != nil {
		t.Error(fmt.Sprintf("SearchQuery failed to return values: %v", err.Error()))
	}

	// check the results and expected results
	expected := "https://scholar.google.co.jp/scholar?hl=en&q=deep+learning+author:\"Bengio\"&as_ylo=2015&as_yhi=&start=20&num=100"
	if query != expected {
		t.Error(fmt.Sprintf("SearchQuery returned unexpected values\n  Expected: %v\n  Query   : %v", expected, query))
	}
	fmt.Printf("SearchQuery() returns %v\n", query)
}

func TestFindQuery(t *testing.T) {
	// exec findQuery()
	query, err := FindQuery("8108748482885444188", "")
	if err != nil {
		t.Error(fmt.Sprintf("findQuery() failed to return values: %v", err.Error()))
	}

	// check the results and expected results
	expected := "https://scholar.google.co.jp/scholar?hl=en&cluster=8108748482885444188&num="
	if query != expected {
		t.Error(fmt.Sprintf("findQuery() returned unexpected values\n  Expected: %v\n  Query   : %v", expected, query))
	}
	fmt.Printf("findQuery() returns %v\n", query)
}

func TestCiteQuery(t *testing.T) {
	// exec citeQuery()
	query, err := CiteQuery("8108748482885444188", "", "2012", "20", "40")
	if err != nil {
		t.Error(fmt.Sprintf("citeQuery() failed to return values: %v", err.Error()))
	}
	expected := "https://scholar.google.co.jp/scholar?hl=en&cites=8108748482885444188&as_ylo=2012&as_yhi=&start=20&num=40"
	// check the results and expected results
	if query != expected {
		t.Error(fmt.Sprintf("citeQuery() returned unexpected values\n  Expected: %v\n  Query   : %v", expected, query))
	}
	fmt.Printf("citeQuery() returns %v\n", query)
}

func TestCitePopQuery(t *testing.T) {
	// set params
	info := "XOJff8gPiHAJ"

	// exec NewQuery()
	query, err := CitePopUpQuery(info)
	if err != nil {
		t.Error("NewQuery() of CitePopUpQuery() failed to return values: ", err.Error())
	}

	// check the results and expected results
	expected := "https://scholar.google.co.jp/scholar?q=info:XOJff8gPiHAJ:scholar.google.com/&output=cite&scirp=0&hl=en"
	if query != expected {
		t.Error(fmt.Sprintf("NewQuery() returned unexpected values\n  Expected: %v\n  Query   : %v", expected, query))
	}
	fmt.Printf("NewQuery() returns %v\n", query)
}