package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/docopt/docopt-go"
	"testing"
	"strconv"
	"errors"
)

func checkWithFirst(query func(map[string]interface{}) (string, error), args []string, aExpected *Article) error {
	arguments, _ := docopt.Parse(USAGE, args[1:], true , VERSION, false)
	doc, err := getDoc(query, arguments)
	if err != nil{
		return err
	}

	a := NewArticle()
	a.Parse(doc.Find(WHOLE_ARTICLE_SELECTOR).First(), false)

	// check
	if !a.same(aExpected) {
		a.showDifference(aExpected)
		return errors.New("") // TODO: fill
	}

	return nil
}

func TestParseCase1(t *testing.T) {
	args := []string{"go-scholar", "find", "3391028632449519147"}

	aExpected := NewArticle()
	aExpected.Title = "Learning with kernels: support vector machines, regularization, optimization, and beyond"
	aExpected.Year = "2002"
	aExpected.URL = "https://books.google.co.jp/books?hl=en&lr=&id=y8ORL3DWt4sC&oi=fnd&pg=PR13&ots=bKyS8zP5Iz&sig=dC5YzrzUAz8kjnEx392vrjb6cr0"
	aExpected.ClusterId = "3391028632449519147"
	aExpected.NumberOfCitations = "10431"
	aExpected.NumberOfVersions = "" // find cannot extract the versions
	aExpected.InfoId = "Kw5VJJNaDy8J"
	aExpected.PDFLink = "" // find cannnot extract the verions
	aExpected.PDFSource = "" // find cannot extract the verions

	if err := checkWithFirst(FindQuery, args, aExpected); err != nil {
		t.Error(err)
	}
}

func TestParseTitle(t *testing.T) {
	// fetch goquery.Document
	url := "https://scholar.google.co.jp/scholar?hl=en&q=\"Learning+deep+architectures+for+AI\"&as_ylo=&as_yhi=&start=&num="
	doc, err := goquery.NewDocument(url)
	if err != nil {
		t.Error(fmt.Sprintf("Fail to get goquery.Document: %v", err.Error()))
	}

	// set expected and actual
	a := NewArticle()
	a.parseTitle(doc.Find(WHOLE_ARTICLE_SELECTOR).First())
	expected := "Learning deep architectures for AI"

	// check
	if a.Title != expected {
		t.Error(fmt.Sprintf("\nExpected: %v\n  Actual: %v", expected, a.Title))
	}
}

func TestParseHeader(t *testing.T) {
	// fetch goquery.Document
	url := "https://scholar.google.co.jp/scholar?hl=en&q=\"Learning+deep+architectures+for+AI\"&as_ylo=&as_yhi=&start=&num="
	doc, err := goquery.NewDocument(url)
	if err != nil {
		t.Error(fmt.Sprintf("failed to get goquery.Document: %v", err.Error()))
	}

	// set expected and actual
	a := NewArticle()
	a.parseHeader(doc.Find(WHOLE_ARTICLE_SELECTOR).First())
	expected := "2009"

	// check
	if a.Year != expected {
		t.Error(fmt.Sprintf("\nExpected: %v\n  Actual: %v", expected, a.Year))
	}
}

func TestParseFooter(t *testing.T) {
	// fetch goquery.Document
	url := "https://scholar.google.co.jp/scholar?hl=en&q=\"Learning+deep+architectures+for+AI\"&as_ylo=&as_yhi=&start=&num="
	doc, err := goquery.NewDocument(url)
	if err != nil {
		t.Error(fmt.Sprintf("failed to get goquery.Document: %v", err.Error()))
	}

	// set expected and actual
	a := NewArticle()
	a.parseFooter(doc.Find(WHOLE_ARTICLE_SELECTOR).First())
	expectedClusterId := "5331804836605365413"
	expectedLowerNumberOfCitations := 2000
	expectedLowerNumberOfVersions := 50
	expectedUpperNumberOfVersions := 100
	expectedInfoId := "pYyS8T9g_kkJ"

	// check
	if a.ClusterId != expectedClusterId {
		t.Error(fmt.Sprintf("\nExpected: %v\n  Actual: %v", expectedClusterId, a.ClusterId))
	}
	c, err := strconv.Atoi(a.NumberOfCitations)
	if err != nil {
		t.Error(fmt.Sprintf("cannot convert # of citations to int: %v", err.Error()))
	}
	if c <= expectedLowerNumberOfCitations {
		t.Error(fmt.Sprintf("\nExpected (more than): %v\n  Actual: %v", expectedLowerNumberOfCitations, a.NumberOfCitations))
	}
	v, err := strconv.Atoi(a.NumberOfVersions)
	if err != nil{
		t.Error(fmt.Sprintf("cannot convert # of versions to int: %v", err.Error()))
	}
	if v <= expectedLowerNumberOfVersions || v >= expectedUpperNumberOfVersions {
		t.Error(fmt.Sprintf("\nExpected (between): %v and %v\n  Actual: %v", expectedLowerNumberOfVersions, expectedUpperNumberOfVersions, v))
	}
	if a.InfoId != expectedInfoId {
		t.Error(fmt.Sprintf("\nExpected: %v\n  Actual: %v", expectedInfoId, a.InfoId))
	}
}

func TestIsValid(t *testing.T) {
	// fetch goquery.Document
	url := "https://scholar.google.co.jp/scholar?hl=en&q=author:\"Bengio\""
	doc, err := goquery.NewDocument(url)
	if err != nil {
		t.Error(fmt.Sprintf("failed to get goquery.Document: %v", err.Error()))
	}

	// parse
	a := NewArticle()
	a.Parse(doc.Find(WHOLE_ARTICLE_SELECTOR).First(), false)

	// check
	expected := false
	if v := a.isValid(); v != expected {
		t.Error(fmt.Sprintf("\nExpected: %v\n  Actual: %v", expected, v))
	}
}

func TestSame(t *testing.T) {

}