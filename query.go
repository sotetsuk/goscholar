package main

import (
	"fmt"
	"strings"
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

func SearchQuery(query, author, title, after, before, start, num string) (string, error) {
	// TODO: validate inputs
	q := query
	if author != "" {
		if !StartAndEndWithDoubleQuotation(author){
			q += "+author:\"" + author + "\""
		}
	}
	if title != "" {
		if !StartAndEndWithDoubleQuotation(title) {
			q += "+\"" + title + "\""
		}
	}
	q = strings.Replace(q, " ", "+", -1)

	return fmt.Sprintf(SEARCH_URL, q, after, before, start, num), nil
}

func FindQuery(cluster_id, num string) (string, error) {
	// TODO: validate inputs
	return fmt.Sprintf(FIND_URL, cluster_id, num), nil
}

func CiteQuery(cluster_id, before, after, start, num string) (string, error) {
	// TODO: validate inputs
	return fmt.Sprintf(CITE_URL, cluster_id, after, before, start, num), nil
}

func CitePopUpQuery(info string) (string, error) {
	// TODO: validate inputs
	return fmt.Sprintf(CITE_POPUP_URL, info), nil
}