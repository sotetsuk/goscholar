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
  go-scholar cite <cluster-id> [search-options] [output-options]
  go-scholar find <cluster-id> [--num=<num>] [output-options]
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
  --num-search=<num-search>
  --start=<start>

Output-options:
  --json
  --bibtex

Others:
  -h --help
  --version`

	arguments, _ := docopt.Parse(usage, os.Args[1:], true, version, false)
	Args := make(map[string]string)

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
	} else if arguments["find"].(bool) {
		// set parameters
		cluster_id := arguments["<cluster-id>"].(string)

		num := ""
		if arguments["--num"] != nil {
			num = arguments["--num"].(string)
		}

		n := 1
		if v, err := strconv.Atoi(num); num != "" && err != nil { // TODO: fix error handling
			n = v
		}

		// get doc
		query, err := FindQuery(cluster_id, num)
		if err != nil {
			log.Fatal(fmt.Sprintf("failed to parse query for find subcommand: %v", err.Error()))
		}

		// parse
		doc, err := goquery.NewDocument(query)
		as := NewArticles(n) // TODO: this n is not appropriate. another parameter is required.
		as.ParseAllArticles(doc, false)

		// output
		fmt.Println(as.Json()) // TODO: check --json, --bibtex
	} else if arguments["cite"].(bool) {

	} else {
		log.Fatal("Wrong arguments.")
	}
}
