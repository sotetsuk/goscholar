[![GoDoc](https://godoc.org/github.com/sotetsuk/goscholar?status.svg)](https://godoc.org/github.com/sotetsuk/goscholar)
[![Build Status](https://travis-ci.org/sotetsuk/goscholar.svg?branch=master)](https://travis-ci.org/sotetsuk/goscholar)
[![Coverage Status](https://coveralls.io/repos/github/sotetsuk/goscholar/badge.svg?branch=master)](https://coveralls.io/github/sotetsuk/goscholar?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/sotetsuk/goscholar)](https://goreportcard.com/report/github.com/sotetsuk/goscholar)
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

### Build
Also, you can use ```build``` command to  build command line tool from the source code.

```
$ git clone git@github.com:sotetsuk/goscholar.git
$ goscholar/build
```

Options:

```
--dev: apply go fmt to all files and save dependencies
```

After ```build``` command executed, you will find corss-compiled binary files in ```bin``` directory.

## Feature

- API for Go 
- API for command line
- search by keywords, title, and author
- find by ```<cluster-id>```
- search the articles citing ```<cluster-id>```
- JSON output
- recursive crawling is not implemented

## Go API

### Example

```go
// create Query and generate URL
q := Query{Keywords:"nature 2015", Author:"y bengio", Title:"Deep learning"}
url := q.SearchUrl()

// fetch document sending the request to the URL
doc, err := Fetch(url)
if err != nil {
	log.Error(err)
	return
}

// parse articles
ch := make(chan *Article, 10)
go ParseDocument(ch, doc)
for a := range ch {
	fmt.Println("---")
	fmt.Println(a)
}
```

## Command line API

### Example

```sh
$ goscholar search --keywords "deep learning nature" --author "y bengio" --after 2015 --num 1 | jq .
[
  {
    "title": {
      "name": "Deep learning",
      "url": "http://www.nature.com/nature/journal/v521/n7553/abs/nature14539.html"
    },
    "year": "2015",
    "cluster_id": "5362332738201102290",
    "num_cite": "499",
    "num_ver": "7",
    "info_id": "0qfs6zbVakoJ",
    "link": {
      "name": "psu.edu",
      "url": "http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.436.894&rep=rep1&type=pdf",
      "format": "PDF"
    },
    "bibtex": "@article{lecun2015deep, title={Deep learning}, author={LeCun, Yann and Bengio, Yoshua and Hinton, Geoffrey}, journal={Nature}, volume={521}, number={7553}, pages={436--444}, year={2015}, publisher={Nature Publishing Group}}",
    "author": [
      "LeCun, Yann",
      "Bengio, Yoshua",
      "Hinton, Geoffrey"
    ],
    "journal": "Nature",
    "booktitle": "",
    "volume": "521",
    "number": "7553",
    "pages": "436--444",
    "publisher": "Nature Publishing Group"
  }
]
```

```sh
$ goscholar find 15502119379559163003 | jq .
[
  {
    "title": {
      "name": "Deep learning via Hessian-free optimization",
      "url": "http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf"
    },
    "year": "2010",
    "cluster_id": "15502119379559163003",
    "num_cite": "269",
    "num_ver": "",
    "info_id": "e6RSJHGXItcJ",
    "link": {
      "name": "wustl.edu",
      "url": "http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf",
      "format": "PDF"
    },
    "bibtex": "@inproceedings{martens2010deep, title={Deep learning via Hessian-free optimization}, author={Martens, James}, booktitle={Proceedings of the 27th International Conference on Machine Learning (ICML-10)}, pages={735--742}, year={2010}}",
    "author": [
      "Martens, James"
    ],
    "journal": "",
    "booktitle": "Proceedings of the 27th International Conference on Machine Learning (ICML-10)",
    "volume": "",
    "number": "",
    "pages": "735--742",
    "publisher": ""
  }
]
```

```sh
$ goscholar cite 15502119379559163003 --num 1 | python -mjson.tool
[
  {
    "title": {
      "name": "Deep neural networks for acoustic modeling in speech recognition: The shared views of four research groups",
      "url": "http://ieeexplore.ieee.org/xpls/abs_all.jsp?arnumber=6296526"
    },
    "year": "2012",
    "cluster_id": "3674494786452480182",
    "num_cite": "1559",
    "num_ver": "27",
    "info_id": "tmCGO4pt_jIJ",
    "link": {
      "name": "toronto.edu",
      "url": "http://www.cs.toronto.edu/~asamir/papers/SPM_DNN_12.pdf",
      "format": "PDF"
    },
    "bibtex": "@article{hinton2012deep, title={Deep neural networks for acoustic modeling in speech recognition: The shared views of four research groups}, author={Hinton, Geoffrey and Deng, Li and Yu, Dong and Dahl, George E and Mohamed, Abdel-rahman and Jaitly, Navdeep and Senior, Andrew and Vanhoucke, Vincent and Nguyen, Patrick and Sainath, Tara N and others}, journal={Signal Processing Magazine, IEEE}, volume={29}, number={6}, pages={82--97}, year={2012}, publisher={IEEE}}",
    "author": [
      "Hinton, Geoffrey",
      "Deng, Li",
      "Yu, Dong",
      "Dahl, George E",
      "Mohamed, Abdel-rahman",
      "Jaitly, Navdeep",
      "Senior, Andrew",
      "Vanhoucke, Vincent",
      "Nguyen, Patrick",
      "Sainath, Tara N",
      "others"
    ],
    "journal": "Signal Processing Magazine, IEEE",
    "booktitle": "",
    "volume": "29",
    "number": "6",
    "pages": "82--97",
    "publisher": "IEEE"
  }
]
```

(This article cites 15502119379559163003=Deep learning via Hessian-free optimization)

### Usage

```
goscholar: Google Scholar crawler and scraper written in Go

Usage:
  goscholar search [--keywords=<keywords>] [--author=<author>] [--title=<title>]
                   [--after=<year>] [--before=<year>] [--num=<num>] [--start=<start>]
                   [--user-agent=<user-agent>]
  goscholar find <cluster-id> [--user-agent=<user-agent>]
  goscholar cite <cluster-id> [--after=<year>] [--before=<year>] [--num=<num>] [--start=<start>]
                              [--user-agent=<user-agent>]
  goscholar -h | --help
  goscholar --version

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
- [github.com/sotetsuk/gobibtex](https://github.com/sotetsuk/gobibtex)

## Related Work
goscholar is inspired by [scholar.py](https://github.com/ckreibich/scholar.py)

## Contribute
Contritubing is more than welcome! See [Issues](https://github.com/sotetsuk/goscholar/issues) for what is required.

## License
MIT License
