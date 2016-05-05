package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/docopt/docopt-go"
	"testing"
	"strconv"
)

type FailParserTestError struct {
	a, aExpected *Article
}

func (e FailParserTestError) Error() string {
	return "\n" + e.a.String() + "\n---\n" + e.aExpected.String()
}

func CheckParseResults(args []string, aExpected *Article) error {
	arguments, _ := docopt.Parse(USAGE, args[1:], true , VERSION, false)
	doc, err := getDoc(FindQuery, arguments)
	if err != nil{
		return nil
	}

	a := NewArticle()
	a.Parse(doc.Find(WHOLE_ARTICLE_SELECTOR).First(), false)

	// check
	if !a.Same(aExpected) {
		return FailParserTestError{a, aExpected}
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
	aExpected.NumberOfVersions = "14"
	aExpected.InfoId = "Kw5VJJNaDy8J"
	aExpected.PDFLink = "http://dip.sun.ac.za/~hanno/tw796/lesings/mlss06au_scholkopf_lk.pdf"
	aExpected.PDFSource = "sun.ac.za"

	if err := CheckParseResults(args, aExpected); err != nil {
		t.Error(err)
	}
}

func TestParseTitle(t *testing.T) {
	// fetch goquery.Document
	url := "https://scholar.google.co.jp/scholar?hl=en&q=\"Learning+deep+architectures+for+AI\"&as_ylo=&as_yhi=&start=&num="
	doc, err := goquery.NewDocument(url)
	if err != nil {
		t.Error(fmt.Sprintf("URL is not valid: %v", err.Error()))
	}

	// set expected and actual
	a := NewArticle()
	a.parseTitle(doc.Find(WHOLE_ARTICLE_SELECTOR).First())
	expected := "Learning deep architectures for AI"

	// check
	if a.Title != expected {
		t.Error(fmt.Sprintf("Expected: %v\nActual: %v", expected, a.Title))
	}
	fmt.Println("Title: ", a.Title)
}

func TestParseHeader(t *testing.T) {
	// fetch goquery.Document
	url := "https://scholar.google.co.jp/scholar?hl=en&q=\"Learning+deep+architectures+for+AI\"&as_ylo=&as_yhi=&start=&num="
	doc, err := goquery.NewDocument(url)
	if err != nil {
		t.Error(fmt.Sprintf("URL is not valid: %v", err.Error()))
	}

	// set expected and actual
	a := NewArticle()
	a.parseHeader(doc.Find(WHOLE_ARTICLE_SELECTOR).First())
	expected := "2009"

	// check
	if a.Year != expected {
		t.Error(fmt.Sprintf("Expected: %v\nActual: %v", expected, a.Year))
	}
	fmt.Println("Year: ", a.Year)
}

func TestParseFooter(t *testing.T) {
	// fetch goquery.Document
	url := "https://scholar.google.co.jp/scholar?hl=en&q=\"Learning+deep+architectures+for+AI\"&as_ylo=&as_yhi=&start=&num="
	doc, err := goquery.NewDocument(url)
	if err != nil {
		t.Error(fmt.Sprintf("URL is not valid: %v", err.Error()))
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
		t.Error(fmt.Sprintf("Expected: %v\nActual: %v", expectedClusterId, a.ClusterId))
	}
	c, err := strconv.Atoi(a.NumberOfCitations)
	if err != nil {
		t.Error(fmt.Sprintf("cannot convert # of citations to int: %v", err.Error()))
	}
	if c <= expectedLowerNumberOfCitations {
		t.Error(fmt.Sprintf("Expected (more than): %v\nActual: %v", expectedLowerNumberOfCitations, a.NumberOfCitations))
	}
	v, err := strconv.Atoi(a.NumberOfVersions)
	if err != nil{
		t.Error(fmt.Sprintf("cannot convert # of versions to int: %v", err.Error()))
	}
	if v <= expectedLowerNumberOfVersions || v >= expectedUpperNumberOfVersions {
		t.Error(fmt.Sprintf("Expected (between): %v and %v\nActual: %v", expectedLowerNumberOfVersions, expectedUpperNumberOfVersions, v))
	}
	if a.InfoId != expectedInfoId {
		t.Error(fmt.Sprintf("Expected: %v\nActual: %v", expectedInfoId, a.InfoId))
	}
	fmt.Println("ClusterId: ", a.ClusterId)
	fmt.Println("NumberOfCitations: : ", a.NumberOfCitations)
	fmt.Println("NumberOfValidations: ", a.NumberOfVersions)
	fmt.Println("InfoId: ", a.InfoId)
}

func TestIsValid(t *testing.T) {
	// fetch goquery.Document
	url := "https://scholar.google.co.jp/scholar?hl=en&q=author:\"Bengio\""
	doc, err := goquery.NewDocument(url)
	if err != nil {
		t.Error(fmt.Sprintf("URL is not valid: %v", err.Error()))
	}

	// parse
	a := NewArticle()
	a.Parse(doc.Find(WHOLE_ARTICLE_SELECTOR).First(), false)

	// check
	expected := false
	if v := a.IsValid(); v != expected {
		t.Error(fmt.Sprintf("Expected: %v\n Actual: %v", expected, v))
	}
}

func TestSame(t *testing.T) {

}