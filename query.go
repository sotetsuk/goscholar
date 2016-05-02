package go_scholar

import (
	"fmt"
	"errors"
)

const (
	SCHOLAR_URL = "https://scholar.google.co.jp/"
)

var (
	SEARCH_URL     = SCHOLAR_URL + "scholar?hl=en&q=%v&as_ylo=%v&as_yhi=%v&start=%v&num=%v"
	FIND_URL       = SCHOLAR_URL + "scholar?hl=en&cluster=%v&num=%v"
	CITE_URL       = SCHOLAR_URL + "scholar?hl=en&cites=%v&as_ylo=%v&as_yhi=%v&start=%v&num=%v"
	CITE_POPUP_URL = SCHOLAR_URL + "scholar?q=info:%s:scholar.google.com/&output=cite&scirp=0&hl=en"
)

type ScholarQuery struct {
	method     string
	parameters map[string]string
}

type CitePopUpQuery struct {
	info string
}

func (q *ScholarQuery) NewQuery() (string, error) {
	if q.method == "search" {

	} else if q.method == "find" {
		return q.findQuery()
	} else if q.method == "cite" {
		return q.citeQuery()
	}

	return "", errors.New("Wrong method was called. Only search, find and cite are usable.")
}

func (q *ScholarQuery) searchQuery() (string, error) {
	return "", nil
}

func (q *ScholarQuery) findQuery() (string, error) {
	// parse parameters
	cluster_id, exists := q.parameters["cluster_id"]
	if !exists {
		return "", errors.New("Parameter cluster_id is needed!")
	}

	// fill nil of option by blank string
	num, exists := q.parameters["num"]
	if !exists {
		num = ""
	}

	return fmt.Sprintf(FIND_URL, cluster_id, num), nil
}

func (q *ScholarQuery) citeQuery() (string, error) {
	// parse parameters
	cluster_id, ok := q.parameters["cluster_id"]
	if !ok {
		return "", errors.New("Parameter cluster_id is needed!")
	}

	// fill nil options by blank string
	params := []string{"before", "after", "num", "start"}
	for _, op := range params {
		if _, exists := q.parameters[op]; !exists {
			q.parameters[op] = ""
		}
	}

	return fmt.Sprintf(CITE_URL, cluster_id, q.parameters["after"], q.parameters["before"], q.parameters["start"], q.parameters["num"]), nil
}

func (q *CitePopUpQuery) NewQuery() (string, error) {
	if q.info == "" {
		return "", errors.New("info attribute is nil!")
	}
	return fmt.Sprintf(CITE_POPUP_URL, q.info), nil
}