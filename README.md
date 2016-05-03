[![Build Status](https://travis-ci.org/sotetsuk/go-scholar.svg?branch=master)](https://travis-ci.org/sotetsuk/go-scholar)

# go-scholar
**Go** ogle **scholar** crawler and scraper written in **Go**


## Install

Assume that `$GOPATH` is set and `$PATH` includes `$GOPATH/bin`.

```
$ go get github.com/sotetsuk/go-scholar
$ go-scholar -h
```

## Usage

```
$ go-scholar -h
go-scholar: scraping google scholar searching results

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
  --num=<num>
  --start=<start>

Output-options:
  --json
  --bibtex

Others:
  -h --help
  --version
```
