package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

type Query struct {
	N int
	clusterId string
	query string
}

func (q *Query) setClusterIdQuery() {
	q.query = SCHOLAR_URL + "scholar?cluster=" + q.clusterId + "&num=" + fmt.Sprintf("%d", q.N)
}

func (q *Query) NewQuery() (*goquery.Document, error) {
	return goquery.NewDocument(q.query)
}