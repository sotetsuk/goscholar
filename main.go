package main

import (
	"github.com/docopt/docopt-go"
	"os"
	"fmt"
	"log"
)

func main() {
	version := "go-scholar 0.0.1"
	usage := `go-scholar: scraping google scholar searching results

Usage:
  go-scholar search [--author=<author>] [--title=<title>] [--query=<query>] [options]
  go-scholar find <cluster-id> [options]
  go-scholar cite <cluster-id> [options]
  go-scholar -h | --help
  go-scholar --version
Options:
  --author=<author>
  --title=<title>
  --query=<query>
  --before=<year>
  --after=<year>
  --num-articles=<num-articles>
  --start=<start>
  --json
  --bibtex
  -h --help
  --version`

	arguments, _ := docopt.Parse(usage, os.Args[1:], true, version, false)
	Args := make(map[string]string)
	options := []string{"--author", "--title", "--query"}

	for _, op := range options {
		if arguments[op] != nil {
			Args[op] = arguments[op].(string)
		}
	}

	if arguments["search"].(bool) {
		fmt.Println("serach")
		ok := false
		for _, op := range options {
			if arguments[op] != nil {
				ok = true
				break
			}
		}
		if !ok {
			log.Fatal("Wrong arguments: at least one of --author, --title or --query is needed.")
		}
	} else if arguments["find"].(bool) {
		fmt.Println("find")
	} else if arguments["cite"].(bool) {
		fmt.Println("cite")
	} else {
		log.Fatal("Wrong arguments.")
	}

	fmt.Println(Args)
}
