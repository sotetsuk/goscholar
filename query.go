package go_scholar

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strconv"
	"log"
)

type Query struct {
	N int
	NCiting int
	ClusterId string
	query string
}

func (q *Query) parseQuery(args map[string]interface{}) {
	var casted = make(map[string]string)
	for key, val := range args {
		casted[key], _ = val.(string)
	}

	q.ClusterId = casted["-C"]
	n, err := strconv.Atoi(casted["-N"])
	if err != nil {
		log.Fatal(err)
	}
	q.N = n
	nCiting, err := strconv.Atoi(casted["-M"])
	if err != nil {
		log.Fatal(err)
	}
	q.NCiting = nCiting
}

func (q *Query) setClusterIdQuery() {
	q.query = SCHOLAR_URL + "scholar?hl=en&cluster=" + q.ClusterId + "&num=" + fmt.Sprintf("%d", q.N)
}

func (q *Query) NewQuery() (*goquery.Document, error) {
	return goquery.NewDocument(q.query)
}