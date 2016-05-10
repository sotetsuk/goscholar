package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/docopt/docopt-go"
	"github.com/sotetsuk/goscholar"
	"os"
)

const articles_buffer = 100000

func main() {
	version := "go-scholar 0.1.2"
	usage := `go-scholar: Google Scholar crawler and scraper written in Go

Usage:
  go-scholar search [--keywords=<keywords>] [--author=<author>] [--title=<title>]
                    [--after=<year>] [--before=<year>] [--num=<num>] [--start=<start>]
  go-scholar find <cluster-id>
  go-scholar cite <cluster-id> [--after=<year>] [--before=<year>] [--num=<num>] [--start=<start>]
  go-scholar -h | --help
  go-scholar --version

Query-options:
  <cluster-id>
  --keywords=<keywords>
  --author=<author>
  --title=<title>

Search-options:
  --after=<year>
  --before=<year>
  --num=<num>
  --start=<start>

Others:
  -h --help
  --version`

	// parse arguments
	args, err := docopt.Parse(usage, os.Args[1:], true, version, false)
	if err != nil {
		log.Error("Failed to parse Args")
		return
	}

	// create Query from arguments
	q := parseArgs(args)
	var url string

	// issue appropriate URL
	if args["search"].(bool) {
		url = q.SearchUrl()
	} else if args["find"].(bool) {
		url = q.FindUrl()
	} else if args["cite"].(bool) {
		url = q.CiteUrl()
	}

	// fetch document from sending the request to the URL
	doc, err := goscholar.Fetch(url)
	if err != nil {
		log.Error(err)
		return
	}

	// parse articles and write
	ch := make(chan *goscholar.Article, articles_buffer)
	go goscholar.ParseDocument(ch, doc)

	fmt.Print("[")
	f := false
	for a := range ch {
		if f {
			fmt.Print(",") // first one is skipped
		}

		fmt.Print(a.Json())
		f = true
	}
	fmt.Print("]")
}

func parseArgs(args map[string]interface{}) (q *goscholar.Query) {
	var keywords, author, title, clusterId, after, before, num, start string

	if args["--keywords"] != nil {
		keywords = args["--keywords"].(string)
	}
	if args["--author"] != nil {
		author = args["--author"].(string)
	}
	if args["--title"] != nil {
		title = args["--title"].(string)
	}
	if args["<cluster-id>"] != nil {
		clusterId = args["<cluster-id>"].(string)
	}
	if args["--after"] != nil {
		after = args["--after"].(string)
	}
	if args["--before"] != nil {
		before = args["--before"].(string)
	}
	if args["--num"] != nil {
		num = args["--num"].(string)
	}
	if args["--start"] != nil {
		start = args["--start"].(string)
	}

	if num == "" {
		num = "10" // as default
	}

	q = &goscholar.Query{
		Keywords:  keywords,
		Author:    author,
		Title:     title,
		ClusterId: clusterId,
		After:     after,
		Before:    before,
		Num:       num,
		Start:     start,
	}

	return q
}
