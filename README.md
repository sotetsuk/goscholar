[![GoDoc](https://godoc.org/github.com/sotetsuk/goscholar?status.svg)](https://godoc.org/github.com/sotetsuk/goscholar)
[![Build Status](https://travis-ci.org/sotetsuk/goscholar.svg?branch=master)](https://travis-ci.org/sotetsuk/goscholar)
[![Coverage Status](https://coveralls.io/repos/github/sotetsuk/goscholar/badge.svg?branch=master)](https://coveralls.io/github/sotetsuk/goscholar?branch=master)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)]()
[![GitHub version](https://badge.fury.io/gh/sotetsuk%2Fgoscholar.svg)](https://badge.fury.io/gh/sotetsuk%2Fgoscholar)

# goscholar
**Go**ogle **Scholar** scraper written in **Go**


## Install

```sh
$ go get github.com/sotetsuk/goscholar
```

for command line:
 
```sh
$ go get github.com/sotetsuk/goscholar/cmd/goscholar
$ goscholar -h
```

## Feature

- API for Go 
- API for command line
- search by keywords, title, and author
- find by ```<cluster-id>```
- search the articles citing ```<cluster-id>```
- JSON output

## Go API

### Example

```go
// create Query and generate URL
q := Query{Keywords:"deep learning", Author:"y bengio"} 
url = q.SearchUrl()

// fetch document sending the request to the URL
doc, err := goscholar.Fetch(url)
if err != nil {
    log.Error(err)
	return
}

// parse articles
ch := make(chan *goscholar.Article, 10)
go goscholar.ParseDocument(ch, doc)

for a := range ch {
	fmt.Println(a)
}
```

## Command line API

### Example

```sh
$ goscholar search --keywords "deep learning nature" --author "y bengio" --after 2015 --num 1 | python -mjson.tool
[
    {
        "ClusterId": "5362332738201102290",
        "InfoId": "0qfs6zbVakoJ",
        "Link": {
            "Format": "PDF",
            "Name": "psu.edu",
            "Url": "http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.436.894&rep=rep1&type=pdf"
        },
        "NumCite": "390",
        "NumVer": "7",
        "Title": {
            "Name": "Deep learning",
            "Url": "http://www.nature.com/nature/journal/v521/n7553/abs/nature14539.html"
        },
        "Year": "2015"
    }
] 
```

```sh
$ goscholar find 15502119379559163003 | python  -mjson.tool
[
    {
        "ClusterId": "15502119379559163003",
        "InfoId": "e6RSJHGXItcJ",
        "Link": {
            "Format": "PDF",
            "Name": "wustl.edu",
            "Url": "http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf"
        },
        "NumCite": "260",
        "NumVer": "",
        "Title": {
            "Name": "Deep learning via Hessian-free optimization",
            "Url": "http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf"
        },
        "Year": "2010"
    }
] 
```

```sh
$ goscholar cite 15502119379559163003 --num 1 | python -mjson.tool
[
    {
        "ClusterId": "3674494786452480182",
        "InfoId": "tmCGO4pt_jIJ",
        "Link": {
            "Format": "PDF",
            "Name": "toronto.edu",
            "Url": "http://www.cs.toronto.edu/~asamir/papers/SPM_DNN_12.pdf"
        },
        "NumCite": "1452",
        "NumVer": "27",
        "Title": {
            "Name": "Deep neural networks for acoustic modeling in speech recognition: The shared views of four research groups",
            "Url": "http://ieeexplore.ieee.org/xpls/abs_all.jsp?arnumber=6296526"
        },
        "Year": "2012"
    }
]
```

(This article cites 15502119379559163003=Deep learning via Hessian-free optimization)

### Usage

```
go-scholar: Google Scholar crawler and scraper written in Go

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
  --version
```

## Dependencies

- [github.com/docopt/docopt-go](https://github.com/docopt/docopt-go)
- [github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)
- [github.com/Sirupsen/logrus](https://github.com/PuerkitoBio/goquery)

## Related Work
goscholar is inspired by [scholar.py](https://github.com/ckreibich/scholar.py)

## Contribute
Contritubing is more than welcome! See [Issues](https://github.com/sotetsuk/goscholar/issues) for what is required.

## License
MIT License
