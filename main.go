package go_scholar

import (
	"github.com/docopt/docopt-go"
	"os"
	"fmt"
	"log"
)

func main() {
	version := "arxiv-api 0.0.1"
	usage := `go-scholar
Usage:
  go-scholar [-C=<cluster-id>] [-N=<num-articles>] [-M=<num-citing>]
  go-scholar -h | --help
  go-scholar --version
Options:
  -C=<cluster-id>        ClusterId
  -N=<num-articles>      Number of articles which we will fetch [default: 1]
  -M=<num-citing>        Number of articles which cites each article [default: 3]
  -h --help              Show this screen.
  --version              Show version.`


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
