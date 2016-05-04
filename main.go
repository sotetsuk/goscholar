package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/docopt/docopt-go"
	"log"
	"os"
)

const ARTICLES_BUFFER = 100000

func main() {
	version := "go-scholar 0.0.1"
	usage := `go-scholar: Google Scholar crawler and scraper written in Go

Usage:
  go-scholar search [--author=<author>] [--title=<title>] [--query=<query>]
                    [--after=<year>] [--before=<year>] [--num=<num>] [--start=<start>]
                    [--json|--bibtex]
  go-scholar find <cluster-id> [--num=<num>] [--json|--bibtex]
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

	arguments, _ := docopt.Parse(usage, os.Args[1:], true, version, false)

	// set parameters
	var author, title, query, cluster_id, after, before, num, start string // TODO: change type of a few options to int

	if arguments["--author"] != nil {
		author = arguments["--author"].(string)
	}
	if arguments["--title"] != nil {
		title = arguments["--title"].(string)
	}
	if arguments["query"] != nil {
		query = arguments["query"].(string)
	}
	if arguments["<cluster-id>"] != nil {
		cluster_id = arguments["<cluster-id>"].(string)
	}
	if arguments["--after"] != nil {
		after = arguments["--after"].(string)
	}
	if arguments["--before"] != nil {
		before = arguments["--before"].(string)
	}
	if arguments["--num"] != nil {
		num = arguments["--num"].(string)
	}
	if arguments["--start"] != nil {
		start = arguments["--start"].(string)
	}

	if num == "" {
		num = "10"
	}

	var doc *goquery.Document

	if arguments["search"].(bool) {
		// issue url
		url, err := SearchQuery(query, author, title, after, before, start, num)
		if err != nil {
			log.Fatal(fmt.Sprintf("failed to parse query for find subcommand: %v", err.Error()))
		}

		// get doc from url
		d, err := goquery.NewDocument(url)
		if err != nil {
			log.Fatal(fmt.Sprintf("failed to get goquery.Document from the query: %v", err.Error()))
		}
		doc = d
	} else if arguments["find"].(bool) { // TODO: remove --num parameter (write --num=1 directly in this block)
		// get doc
		url, err := FindQuery(cluster_id, num)
		if err != nil {
			log.Fatal(fmt.Sprintf("failed to parse query for find subcommand: %v", err.Error()))
		}
		d, err := goquery.NewDocument(url)
		if err != nil {
			log.Fatal(fmt.Sprintf("failed to get goquery.Document from the query: %v", err.Error()))
		}
		doc = d
	} else if arguments["cite"].(bool) {
		// get doc
		url, err := CiteQuery(cluster_id, after, before, num, start)
		if err != nil {
			log.Fatal(fmt.Sprintf("failed to parse query for find subcommand: %v", err.Error()))
		}
		d, err := goquery.NewDocument(url)
		if err != nil {
			log.Fatal(fmt.Sprintf("failed to get goquery.Document from the query: %v", err.Error()))
		}
		doc = d
	} else {
		log.Fatal("Wrong arguments. [search|find|cite] is valid.")
	}

	// parse and output
	as := NewArticles(ARTICLES_BUFFER)
	go as.ParseAllArticles(doc, false)
	as.StdoutJson()  // TODO: treat --json|--bibtex parameters
}
