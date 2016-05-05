[![Build Status](https://travis-ci.org/sotetsuk/go-scholar.svg?branch=master)](https://travis-ci.org/sotetsuk/go-scholar)
[![Coverage Status](https://coveralls.io/repos/github/sotetsuk/go-scholar/badge.svg?branch=master)](https://coveralls.io/github/sotetsuk/go-scholar?branch=master)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)]()
[![GitHub version](https://badge.fury.io/gh/sotetsuk%2Fgo-scholar.svg)](https://badge.fury.io/gh/sotetsuk%2Fgo-scholar)

# go-scholar
**Go**ogle **Scholar** crawler and scraper written in **Go**


## Install

Assume that `$GOPATH` is set and `$PATH` includes `$GOPATH/bin`.

```
$ go get github.com/sotetsuk/go-scholar
$ go-scholar -h
```

## Feature
- search by keywords, title, and author
- find by ```<cluster-id>```
- fetch citing articles of ```<cluster-id>```
- crawl recursively (**not implemented yet**)
- JSON output

## Example 

```
$ go-scholar search --title "Deep learning via Hessian-free optimization" --num 1 | python -mjson.tool
[
    {
        "ClusterId": "15502119379559163003",
        "InfoId": "e6RSJHGXItcJ",
        "NumberOfCitations": "260",
        "NumberOfVersions": "9",
        "PDFLink": "http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf",
        "PDFSource": "wustl.edu",
        "Title": "Deep learning via Hessian-free optimization",
        "URL": "http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf",
        "Year": "2010"
    }
]
```


```
$ go-scholar find 8108748482885444188 | python -mjson.tool
[
    {
        "ClusterId": "8108748482885444188",
        "InfoId": "XOJff8gPiHAJ",
        "NumberOfCitations": "376",
        "NumberOfVersions": "",
        "PDFLink": "",
        "PDFSource": "",
        "Title": "Learning in science: A comparison of deep and surface approaches",
        "URL": "http://onlinelibrary.wiley.com/doi/10.1002/(SICI)1098-2736(200002)37:2%3C109::AID-TEA3%3E3.0.CO;2-7/abstract",
        "Year": "2000"
    }
]
```

## Usage

```
$ go-scholar -h
go-scholar: Google Scholar crawler and scraper written in Go

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
  --version
```

## Contribute
Contritubing is more than welcome! See [Issues](https://github.com/sotetsuk/go-scholar/issues) for what is required.

## License
MIT License
