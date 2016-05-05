package main

import (
	"fmt"
	"strings"
)

const (
	SCHOLAR_URL = "https://scholar.google.co.jp/"
)

var (
	SEARCH_URL     = SCHOLAR_URL + "scholar?hl=en&q=%v&as_ylo=%v&as_yhi=%v&num=%v&start=%v"
	FIND_URL       = SCHOLAR_URL + "scholar?hl=en&cluster=%v&num=1"
	CITE_URL       = SCHOLAR_URL + "scholar?hl=en&cites=%v&as_ylo=%v&as_yhi=%v&num=%v&start=%v"
	CITE_POPUP_URL = SCHOLAR_URL + "scholar?q=info:%s:scholar.google.com/&output=cite&scirp=0&hl=en"
)

func SearchQuery(arguments map[string]interface{}) (string, error) {
	// TODO: validate inputs
	author, title, query, _, after, before, num, start := parseAndInitializeArguments(arguments)

	searchQuery := func(query, author, title, after, before, num, start string) (string, error) {
		q := query
		if author != "" {
			if !StartAndEndWithDoubleQuotation(author) {
				q += "+author:\"" + author + "\""
			}
		}
		if title != "" {
			if !StartAndEndWithDoubleQuotation(title) {
				q += "+\"" + title + "\""
			}
		}
		q = strings.Replace(q, " ", "+", -1)

		return fmt.Sprintf(SEARCH_URL, q, after, before, num, start), nil
	}
	return searchQuery(query, author, title, after, before, num, start)
}

func FindQuery(arguments map[string]interface{}) (string, error) {
	// TODO: validate inputs
	_, _, _, cluster_id, _, _, _, _ := parseAndInitializeArguments(arguments)

	findQuery := func(cluster_id string) (string, error) {
		return fmt.Sprintf(FIND_URL, cluster_id), nil
	}

	return findQuery(cluster_id)
}

func CiteQuery(arguments map[string]interface{}) (string, error) {
	// TODO: validate inputs
	_, _, _, cluster_id, after, before, num, start := parseAndInitializeArguments(arguments)

	citeQuery := func(cluster_id, after, before, num, start string) (string, error) {
		return fmt.Sprintf(CITE_URL, cluster_id, after, before, num, start), nil
	}

	return citeQuery(cluster_id, after, before, num, start)
}

func CitePopUpQuery(info string) (string, error) {
	// TODO: validate inputs
	return fmt.Sprintf(CITE_POPUP_URL, info), nil
}
