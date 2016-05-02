package go_scholar

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"strconv"
)

type Article struct {
	Title             string
	Year              string
	URL               string
	ClusterId         string
	NumberOfCitations string
	NumberOfVersions  string
	InfoId            string
	PDFLink           string
	PDFSource         string
	Bibtex            string
	CitingId          string
}

type Articles struct {
	n        int
	articles []Article
}

func (a *Article) Parse(s *goquery.Selection, nCiting int, recur bool) {
	a.parseTitle(s)
	a.parseHeader(s)
	a.parseFooter(s)
	a.parseSideBar(s)
	a.parseBibTeX()
	if recur {
		a.parseCitingId(nCiting)
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

func (a *Article) parseBibTeX() {
	popURL := SCHOLAR_URL + "scholar?q=info:" + a.InfoId + ":scholar.google.com/&output=cite&scirp=0&hl=en"
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

func (a *Article) parseCitingId(nCiting int) {
	nCitingStr := strconv.Itoa(nCiting)
	citingURL := SCHOLAR_URL + "scholar?cites=" + a.ClusterId + "&sciodt=0,5&hl=en&num=" + nCitingStr
	citingDoc, err := goquery.NewDocument(citingURL)
	if err != nil {
		log.Fatal(err)
	}
	as := Articles{n:nCiting} // TODO:
	as.ParseAllArticles(citingDoc, 0, false)
	a.CitingId = ""
	for _, b := range as.articles {
		a.CitingId += b.ClusterId + ","
	}
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
	fmt.Println("citingId: ", a.CitingId)
}


func (as *Articles) ParseAllArticles(doc *goquery.Document, nCiting int, recur bool) {
	as.articles = make([]Article, as.n)

	parse := func(i int, s *goquery.Selection) {
		as.articles[i].Parse(s, nCiting, recur)
	}
	doc.Find(WHOLE_ARTICLE_SELECTOR).Each(parse)
}

func (as *Articles) Json() string {
	bytes, err := json.Marshal(as.articles)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}
