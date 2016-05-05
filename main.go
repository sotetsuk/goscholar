package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/docopt/docopt-go"
	log "github.com/Sirupsen/logrus"
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

func init() {
	log.SetOutput(os.Stderr)
	log.SetLevel(log.InfoLevel)
}

func main() {
	arguments, _ := docopt.Parse(USAGE, os.Args[1:], true, VERSION, false) // TODO: change type of a few options to int
	log.WithFields(log.Fields{"arguments": arguments}).Debug("arguments are parsed")

	var doc *goquery.Document

	if arguments["search"].(bool) {
		d, err := getDoc(SearchQuery, arguments)
		if err != nil{
			log.Error("Exit for getDoc's failure")
			return
		}
		doc = d
	} else if arguments["find"].(bool) {
		d, err := getDoc(FindQuery, arguments)
		if err != nil{
			return
		}
		doc = d
	} else if arguments["cite"].(bool) {
		d, err := getDoc(CiteQuery, arguments)
		if err != nil{
			return
		}
		doc = d
	}

	// parse and output
	ch := make(chan *Article, ARTICLES_BUFFER)
	go ParseArticles(ch, doc)
	StdoutArticleAsJson(ch)  // TODO: treat --json|--bibtex parameters
}
