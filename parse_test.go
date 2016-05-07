package goscholar

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"testing"
	"time"
)

var url1, url2 string
var doc1, doc2 *goquery.Document
var err1, err2 error
var a1Expected, a2Expected *Article

func init() {
	// set doc1
	url1 = "https://scholar.google.co.jp/scholar?hl=en&q=\"Learning+deep+architectures+for+AI\"&as_ylo=&as_yhi=&start=&num="
	doc1, err1 = Fetch(url1)

	time.Sleep(3 * time.Second)

	// set doc2
	url2 = "https://scholar.google.co.jp/scholar?q=%22Unsupervised+feature+learning+and+deep+learning%3A+A+review+and+new+perspectives%22&btnG=&hl=en"
	doc2, err2 = Fetch(url2)

	// set a1Expected
	a1Expected = &Article{
		Title: &Title{
			Name: "Learning deep architectures for AI",
			Url:  "http://dl.acm.org/citation.cfm?id=1658424",
		},
		Year:      "2009",
		ClusterId: "5331804836605365413",
		NumCite:   "2424",
		NumVer:    "58",
		InfoId:    "pYyS8T9g_kkJ",
		Link: &Link{
			Name:   "sanghv.com",
			Url:    "http://sanghv.com/download/soft/machine%20learning,%20artificial%20intelligence,%20mathematics%20ebooks/ML/learning%20deep%20architectures%20for%20AI%20(2009).pdf",
			Format: "PDF",
		},
	}

	// set a2Expected
	a2Expected = &Article{
		Title: &Title{
			Name: "Unsupervised feature learning and deep learning: A review and new perspectives",
			Url:  "",
		},
		Year:      "2012",
		ClusterId: "1290425714414184502",
		NumCite:   "140",
		NumVer:    "",
		InfoId:    "NuivrFmD6BEJ",
		Link: &Link{
			Name:   "",
			Url:    "",
			Format: "",
		},
	}
}

func checkWithFirst(doc *goquery.Document, aExpected *Article) (err error) {
	a, err := ParseSelection(doc.Find(WHOLE_ARTICLE_SELECTOR).First())

	if !same(a, aExpected) {
		showDifference(a, aExpected)
		return errors.New("Two articles should be same")
	}

	return nil
}

func ExampleParseDocument() {
	time.Sleep(3 * time.Second)

	ch := make(chan *Article, 10000)

	url := "https://scholar.google.co.jp/scholar?hl=en&q=\"deep+learning\"+author:\"y+bengio\"&as_ylo=2015&num=1"
	doc, err := Fetch(url)
	if err != nil {
		return
	}

	go ParseDocument(ch, doc)
	a := <- ch
	fmt.Println(a)

	return
	// Output:
	// -----------------------------------------------------------------------------
	// [Title]
	//   Name: Deep learning
	//   Url: http://www.nature.com/nature/journal/v521/n7553/abs/nature14539.html
	// [Year]
	//   2015
        // [ClusterId]
	//   5362332738201102290
	// [NumeCite]
	//   390
	// [NumVer]
	//   7
	// [InfoId]
	//   0qfs6zbVakoJ
	// [Link]
	//   Name: psu.edu
	//   Url: http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.436.894&rep=rep1&type=pdf
	//   Format: PDF
}

func TestParseSelection(t *testing.T) {
	if err1 != nil {
		t.Skip(err1)
	}
	if err2 != nil {
		t.Skip(err2)
	}

	var err error
	check := func(doc *goquery.Document, aExpected *Article) {
		if err != nil {
			return
		}
		err = checkWithFirst(doc, aExpected)
	}

	check(doc1, a1Expected)
	check(doc2, a2Expected)

	if err != nil {
		t.Error(err)
	}
}

func TestParseH3(t *testing.T) {
	// test case 1
	if err1 != nil {
		t.Skip(err1)
	}

	actual := parseH3(doc1.Find(WHOLE_ARTICLE_SELECTOR).First())

	if actual.Name != a1Expected.Title.Name {
		t.Error(TestErr{expected: a1Expected.Title.Name, actual: actual.Name})
	}

	if actual.Url != a1Expected.Title.Url {
		t.Error(TestErr{expected: a1Expected.Title.Name, actual: actual.Url})
	}

	// test case 2
	if err2 != nil {
		t.Skip(err2)
	}

	actual = parseH3(doc2.Find(WHOLE_ARTICLE_SELECTOR).First())

	if actual.Name != a2Expected.Title.Name {
		t.Error(TestErr{expected: a2Expected.Title.Name, actual: actual.Name})
	}

	if actual.Url != a2Expected.Title.Url {
		t.Error(TestErr{expected: a2Expected.Title.Url, actual: actual.Url})
	}
}

func TestParseGreenLine(t *testing.T) {
	if err1 != nil {
		t.Skip(err1)
	}

	actual := parseGreenLine(doc1.Find(WHOLE_ARTICLE_SELECTOR).First())

	if actual != a1Expected.Year {
		t.Error(TestErr{expected: a1Expected.Year, actual: actual})
	}
}

func TestParseBottom(t *testing.T) {
	if err1 != nil {
		t.Skip(err1)
	}

	clusterId, numCite, numVer, infoId := parseBottom(doc1.Find(WHOLE_ARTICLE_SELECTOR).First())
	a1ExpectedLowerNimCite := 2000
	a1ExpectedLowerNumVer := 50
	a1ExpectedUpperNumVer := 100

	if clusterId != a1Expected.ClusterId {
		t.Error(TestErr{expected: a1Expected.ClusterId, actual: clusterId})
	}
	c, err := strconv.Atoi(numCite)
	if err != nil {
		t.Error(err)
	}

	if c <= a1ExpectedLowerNimCite {
		t.Error("Lower bound error:", TestErr{expected: strconv.Itoa(a1ExpectedLowerNimCite), actual: numCite})
	}
	v, err := strconv.Atoi(numVer)
	if err != nil {
		t.Error(err)
	}
	if v <= a1ExpectedLowerNumVer {
		t.Error("Lower bound error: ", TestErr{expected: strconv.Itoa(a1ExpectedLowerNumVer), actual: numVer})
	}
	if v >= a1ExpectedUpperNumVer {
		t.Error("Upper bound error: ", TestErr{expected: strconv.Itoa(a1ExpectedLowerNumVer), actual: numVer})
	}

	if infoId != a1Expected.InfoId {
		t.Error(TestErr{expected: a1Expected.InfoId, actual: infoId})
	}
}

func TestParseSideBar(t *testing.T) {
	if err1 != nil {
		t.Skip(err1)
	}

	link := parseSideBar(doc1.Find(WHOLE_ARTICLE_SELECTOR).First())

	if link.Name != a1Expected.Link.Name {
		t.Error(TestErr{expected: a1Expected.Link.Name, actual: link.Name})
	}
	if link.Url != a1Expected.Link.Url {
		t.Error(TestErr{expected: a1Expected.Link.Url, actual: link.Url})
	}
	if link.Format != a1Expected.Link.Format {
		t.Error(TestErr{expected: a1Expected.Link.Format, actual: link.Format})
	}
}
