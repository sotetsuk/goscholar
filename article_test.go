package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"testing"
	"strconv"
)


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