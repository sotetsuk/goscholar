package go_scholar

import (
	"testing"
	"fmt"
)

func TestFindQuery(t *testing.T) {
	// set params
	params := map[string]string{"cluster_id":"8108748482885444188", "num":""}
	q := ScholarQuery{"find", params}

	// exec findQuery()
	query, err := q.findQuery()
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
	// set params
	params := map[string]string{"cluster_id":"8108748482885444188", "after":"2012", "start":"20", "num":"40"}
	q := ScholarQuery{"cite", params}

	// exec citeQuery()
	cite_query, err := q.citeQuery()
	if err != nil {
		t.Error(fmt.Sprintf("citeQuery() failed to return values: %v", err.Error()))
	}
	expected := "https://scholar.google.co.jp/scholar?hl=en&cites=8108748482885444188&as_ylo=2012&as_yhi=&start=20&num=40"
	// check the results and expected results
	if cite_query != expected {
		t.Error(fmt.Sprintf("citeQuery() returned unexpected values\n  Expected: %v\n  Query   : %v", expected, cite_query))
	}
	fmt.Printf("citeQuery() returns %v\n", cite_query)
}