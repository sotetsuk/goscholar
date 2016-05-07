package goscholar

import (
	"strings"
	"fmt"
)

type Query struct {
	keywords string
	author string
	title string
	cluster_id string
	info_id string
	after string
	before string
	num string
	start string
}

// SearchUrl generates URL to which Fetch sends request
func (q *Query) SearchUrl() (url string) {
	k := q.keywords
	if q.author != "" {
		if enclosedInDoubleQuotation(q.author) {
			k += "+author:" + q.author
		} else {
			k += "+author:\"" + q.author + "\""
		}
	}
	if q.title != "" {
		if enclosedInDoubleQuotation(q.title) {
			k += q.title
		} else {
			k += "+\"" + q.title + "\""
		}
	}
	k = strings.Replace(k, " ", "+", -1)

	return fmt.Sprintf(SEARCH_URL, k, q.after, q.before, q.num, q.start)
}

func (q *Query) FindUrl() (url string) {
	return fmt.Sprintf(FIND_URL, q.cluster_id)
}

func (q *Query) CiteUrl() (url string) {
	return fmt.Sprintf(CITE_URL, q.cluster_id, q.after, q.before, q.num, q.start)
}

func (q *Query) CitePopUpQueryUrl() (url string) {
	return fmt.Sprintf(CITE_POPUP_URL, q.info_id)
}
