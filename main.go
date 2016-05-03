package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/docopt/docopt-go"
	"log"
	"os"
	"strconv"
)

func main() {
	version := "go-scholar 0.0.1"
	usage := `go-scholar: scraping google scholar searching results

Usage:
  go-scholar search [--author=<author>] [--title=<title>] [--query=<query>] [search-options] [output-options]
  go-scholar find <cluster-id> [--num=<num>] [output-options]
  go-scholar cite <cluster-id> [--after=<year>] [--before=<year>] [--num=<num>] [--start=<start>] [output-options]
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
	Args := make(map[string]string)

	// set parameters
	var cluster_id, after, before, num, start string // TODO: change type to int
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

	n := 10
	if m, err := strconv.Atoi(num); err != nil {
		log.Fatal(fmt.Sprintf("failed to parse --num. --num parameter should be int: %v", err.Error()))
	} else {
		n = m
	}

	var doc *goquery.Document

	if arguments["search"].(bool) {
		query_options := []string{"--author", "--title", "--query"}
		for _, op := range query_options {
			if arguments[op] != nil {
				Args[op] = arguments[op].(string)
			}
		}

		ok := false
		for _, op := range query_options {
			if arguments[op] != nil {
				ok = true
				break
			}
		}
		if !ok {
			log.Fatal("Wrong arguments: at least one of --author, --title or --query is needed.")
		}
	} else if arguments["find"].(bool) { // TODO: remove --num parameter (write --num=1 directly in this block)
		if arguments["--num"] == nil && n == 10 {
			n = 1
		}

		// get doc
		query, err := FindQuery(cluster_id, num)
		if err != nil {
			log.Fatal(fmt.Sprintf("failed to parse query for find subcommand: %v", err.Error()))
		}
		d, err := goquery.NewDocument(query)
		if err != nil {
			log.Fatal(fmt.Sprintf("failed to get goquery.Document from the query: %v", err.Error()))
		}
		doc = d
	} else if arguments["cite"].(bool) {
		// get doc
		query, err := CiteQuery(cluster_id, after, before, num, start)
		if err != nil {
			log.Fatal(fmt.Sprintf("failed to parse query for find subcommand: %v", err.Error()))
		}
		d, err := goquery.NewDocument(query)
		if err != nil {
			log.Fatal(fmt.Sprintf("failed to get goquery.Document from the query: %v", err.Error()))
		}
		doc = d
	} else {
		log.Fatal("Wrong arguments. [search|find|cite] is valid.")
	}

	// parse
	as := NewArticles(n) // TODO: this n is not appropriate. another parameter is required.
	as.ParseAllArticles(doc, false)

	// results > STDOUT
	fmt.Println(as.Json()) // TODO: check --json, --bibtex
}
