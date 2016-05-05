package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/docopt/docopt-go"
	"log"
	"os"
)

const (
	ARTICLES_BUFFER = 100000
	VERSION         = "go-scholar 0.0.1"
	USAGE           = `go-scholar: Google Scholar crawler and scraper written in Go

Usage:
  go-scholar search [--author=<author>] [--title=<title>] [--query=<query>]
                    [--after=<year>] [--before=<year>] [--num=<num>] [--start=<start>]
                    [--json|--bibtex]
  go-scholar find <cluster-id> [--json|--bibtex]
  go-scholar cite <cluster-id> [--after=<year>] [--before=<year>] [--num=<num>] [--start=<start>] [--json|--bibtex]
  go-scholar -h | --help
  go-scholar --version

Query-options:
  <cluster-id>
  --author=<author>
  --title=<title>
  --query=<query>

Search-options:
  --after=<year>
  --before=<year>
  --num=<num>
  --start=<start>

Output-options:
  --json
  --bibtex

Others:
  -h --help
  --version`
)

func main() {
	arguments, _ := docopt.Parse(USAGE, os.Args[1:], true, VERSION, false) // TODO: change type of a few options to int

	var doc *goquery.Document

	if arguments["search"].(bool) {
		doc, _ = getDoc(SearchQuery, arguments)
	} else if arguments["find"].(bool) {
		doc, _ = getDoc(FindQuery, arguments)
	} else if arguments["cite"].(bool) {
		doc, _ = getDoc(CiteQuery, arguments)
	} else {
		log.Fatal("Wrong arguments. [search|find|cite] is valid.")
	}

	// parse and output
	as := NewArticles(ARTICLES_BUFFER)
	go as.ParseAllArticles(doc, false)
	as.StdoutJson()  // TODO: treat --json|--bibtex parameters
}
