package main

import (
	"github.com/docopt/docopt-go"
	"os"
	"fmt"
)

func main() {
	version := "arxiv-api 0.0.1"
	usage := `go-scholar
Usage:
  go-scholar [-C=<cluster-id>] [-N=<num-articles>]
  go-scholar -h | --help
  go-scholar --version
Options:
  -C=<cluster-id>        ClusterId
  -N=<num-articles>      Number of articles which we will fetch [default: 1]
  -h --help              Show this screen.
  --version              Show version.`

	arguments, _ := docopt.Parse(usage, os.Args[1:], true, version, false)

	fmt.Println(arguments)
	q := Query{}
	q.parseQuery(arguments)
	q.setClusterIdQuery()
	doc, _ := q.NewQuery()
	as := Articles{n:q.N}
	as.ParseAllArticles(doc)
	fmt.Println(as.Json())
}
