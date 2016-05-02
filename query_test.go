package go_scholar

import (
	"testing"
	"fmt"
)

func TestQuery(t *testing.T) {
	// test ScholarQuery
	// test NewQuery()
	// test search query
	// test find query
	find_params := map[string]string{"cluster_id":"8108748482885444188", "num":""}
	find_q := ScholarQuery{"find", find_params}
	find_query, err := find_q.findQuery()
	if err != nil {
		t.Error(fmt.Sprintf("findQuery() failed to return values: %v", err.Error()))
	}
	find_expected := "https://scholar.google.co.jp/scholar?hl=en&cluster=8108748482885444188&num="
	if find_query != find_expected {
		t.Error(fmt.Sprintf("findQuery() returned unexpected values\n  Expected: %v\n  Query   : %v", find_expected, find_query))
	}
	fmt.Printf("findQuery() returns %v\n", find_query)

	// test cite query
	cite_params := map[string]string{"cluster_id":"8108748482885444188", "after":"2012", "start":"20", "num":"40"}
	cite_q := ScholarQuery{"cite", cite_params}
	cite_query, err := cite_q.citeQuery()
	if err != nil {
		t.Error(fmt.Sprintf("citeQuery() failed to return values: %v", err.Error()))
	}
	cite_expected := "https://scholar.google.co.jp/scholar?hl=en&cites=8108748482885444188&as_ylo=2012&as_yhi=&start=20&num=40"
	if cite_query != cite_expected {
		t.Error(fmt.Sprintf("citeQuery() returned unexpected values\n  Expected: %v\n  Query   : %v", cite_expected, cite_query))
	}
	fmt.Printf("citeQuery() returns %v\n", cite_query)

	// test CitePopUpQuery
	// test NewQuery()
}