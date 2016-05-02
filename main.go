package go_scholar

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
  go-scholar search (--author=<author>|--title=<title>|--query=<query>) [--before=<year>|--after=<year>|--num-articles=<num-articles>|--start=<start>]
  go-scholar find <cluster-id> [--before=<year>|--after=<year>|--num-articles=<num-articles>|--start=<start>]
  go-scholar cite <cites-id> [--before=<year>|--after=<year>|--num-articles=<num-articles>|--start=<start>]
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
  -h --help
  --version`

	arguments, _ := docopt.Parse(usage, os.Args[1:], true, version, false)

	q := Query{}
	q.parseQuery(arguments)
	q.setClusterIdQuery()
	doc, err := q.NewQuery()
	if err != nil {
		log.Fatal(err)
	}
	as := Articles{n:q.N}
	as.ParseAllArticles(doc, q.NCiting, true)
	fmt.Println(as.Json())
}
