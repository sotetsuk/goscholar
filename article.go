package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

const (
	ARTICLE_TITLE_SELECTOR   = ".gs_rt > a"
	ARTICLE_HEADER_SELECTOR  = ".gs_a"
	ARTICLE_FOOTER_SELECTOR  = ".gs_fl"
	ARTICLE_SIDEBAR_SELECTOR = ".gs_md_wp > a"
)

type Article struct {
	Title             string
	Year              string
	Authors           []string
	URL               string
	ClusterId         string
	NumberOfCitations string
	NumberOfVersions  string
	InfoId            string
	PDFLink           string
	PDFSource         string
	Bibtex            string
}

func NewArticle() *Article {
	a := Article{}
	return &a
}

func (a *Article) Parse(s *goquery.Selection, useBibTeX bool) {
	a.parseTitle(s)
	a.parseHeader(s)
	a.parseFooter(s)
	a.parseSideBar(s)
	if useBibTeX {
		a.crawlAndParseBibTeX()
	}
}

func (a *Article) parseTitle(s *goquery.Selection) {
	h3Title := s.Find(ARTICLE_TITLE_SELECTOR)
	a.URL, _ = h3Title.Attr("href")
	a.Title = h3Title.Text()
}

func (a *Article) parseHeader(s *goquery.Selection) {
	a.Year = parseYear(s.Find(ARTICLE_HEADER_SELECTOR).Text())
}

func (a *Article) parseFooter(s *goquery.Selection) {
	divFooter := s.Find(ARTICLE_FOOTER_SELECTOR)
	parseFooter := func(i int, s *goquery.Selection) {

		href, _ := s.Attr("href")
		text := s.Text()

		if strings.HasPrefix(href, "/scholar?cites") {
			a.ClusterId = parseClusterId(href) // TODO: 両方で
			a.NumberOfCitations = parseNumberOfCitations(text)
		}
		if strings.HasPrefix(href, "/scholar?cluster") {
			a.NumberOfVersions = parseNumberOfVersions(text)
		}
		if strings.HasPrefix(href, "/scholar?q=related") {
			a.InfoId = parseInfoId(href)
		}

	}
	divFooter.Find("a").Each(parseFooter)
}

func (a *Article) parseSideBar(s *goquery.Selection) {
	sideBarA := s.Find(ARTICLE_SIDEBAR_SELECTOR)
	a.PDFLink, _ = sideBarA.Attr("href")
	a.PDFSource = sideBarA.Text()
}

func (a *Article) crawlAndParseBibTeX() {
	popURL, err := CitePopUpQuery(a.InfoId)
	if err != nil {
		log.Fatal(err)
	}

	popDoc, err := goquery.NewDocument(popURL)
	if err != nil {
		log.Fatal(err)
	}

	bibURL, _ := popDoc.Find("#gs_citi > a:first-child").Attr("href")
	bibDoc, err := goquery.NewDocument(SCHOLAR_URL + bibURL)
	if err != nil {
		log.Fatal(err)
	}
	a.Bibtex = bibDoc.Text()
}

func (a *Article) dump() {
	fmt.Println("title :", a.Title)
	fmt.Println("year :", a.Year)
	fmt.Println("url: ", a.URL)
	fmt.Println("cluster_id: ", a.ClusterId)
	fmt.Println("# of citations: ", a.NumberOfVersions)
	fmt.Println("# of versions: ", a.NumberOfCitations)
	fmt.Println("infor id: ", a.InfoId)
	fmt.Println("pdfLink: ", a.PDFLink)
	fmt.Println("pdfSource: ", a.PDFSource)
	fmt.Println("BibTeX: ", a.Bibtex)
}