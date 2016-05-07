package goscholar

import (
	"strings"
	"fmt"
)

// Query issue an appropriate URL to which Fetch sends a request.
type Query struct {
	Keywords string
	Author string
	Title string
	ClusterId string
	InfoId string
	After string
	Before string
	Num string
	Start string
}

// SearchUrl issues URL whose search query is composed of keywords, author and title.
// SearchUrl uses Keywords, Author, Title, After, Before, Num and Start Attributes.
// For example:
//   https://scholar.google.co.jp/scholar?hl=en&q=deep+learning+author:"y+bengio"&as_ylo=2015&as_yhi=&num=100&start=20
func (q *Query) SearchUrl() (url string) {
	k := q.Keywords
	if q.Author != "" {
		if enclosedInDoubleQuotation(q.Author) {
			k += "+author:" + q.Author
		} else {
			k += "+author:\"" + q.Author + "\""
		}
	}
	if q.Title != "" {
		if enclosedInDoubleQuotation(q.Title) {
			k += q.Title
		} else {
			k += "+\"" + q.Title + "\""
		}
	}
	k = strings.Replace(k, " ", "+", -1)

	return fmt.Sprintf(search_url, k, q.After, q.Before, q.Num, q.Start)
}

// FindUrl uses ClusterId which identify the desired article and spits out URL like:
//   https://scholar.google.co.jp/scholar?hl=en&cluster=5362332738201102290&num=1
// FindUrl depends only on ClusterId
func (q *Query) FindUrl() (url string) {
	return fmt.Sprintf(find_url, q.ClusterId)
}

// CiteUrl uses ClusterId and issues the URL whose results include the articles citing the article of the ClusterId.
// This depends on ClusterId, After, Before, Num and Start. For example:
//   https://scholar.google.co.jp/scholar?hl=en&cites=5362332738201102290&as_ylo=2012&as_yhi=&num=40&start=20
func (q *Query) CiteUrl() (url string) {
	return fmt.Sprintf(cite_url, q.ClusterId, q.After, q.Before, q.Num, q.Start)
}

// CitePopUpQueryUrl issues a URL used for getting BibTeX information. This only uses the InfoId. For example:
//   https://scholar.google.co.jp/scholar?q=info:XOJff8gPiHAJ:scholar.google.com/&output=cite&scirp=0&hl=en
func (q *Query) CitePopUpQueryUrl() (url string) {
	return fmt.Sprintf(cite_popup_url, q.InfoId)
}
