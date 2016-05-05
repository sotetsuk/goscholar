package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"regexp"
)

const (
	ARTICLE_TITLE_SELECTOR   = "h3.gs_rt > a"
	ARTICLE_HEADER_SELECTOR  = ".gs_a"
	ARTICLE_FOOTER_SELECTOR  = ".gs_fl"
	ARTICLE_SIDEBAR_SELECTOR = ".gs_md_wp > a"
)

type Article struct {
	Title             string
	Year              string
	// Authors           []string
	URL               string
	ClusterId         string
	NumberOfCitations string
	NumberOfVersions  string
	InfoId            string
	PDFLink           string
	PDFSource         string
	// Bibtex            string
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
	/*
	if useBibTeX {
		a.crawlAndParseBibTeX()
	}
	*/
}

func (a *Article) parseTitle(s *goquery.Selection) {
	h3Title := s.Find(ARTICLE_TITLE_SELECTOR)
	url, exists := h3Title.Attr("href")

	if exists {
		a.URL = url
		a.Title = h3Title.Text()
	} else {
		title := s.Find("h3").Text()
		rep, _ := regexp.Compile("\\[[a-zA-Z0-9]*\\]\\[[a-zA-Z0-9]*\\]\\s")
		a.Title = rep.ReplaceAllString(title, "")
	}
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
			a.ClusterId = parseClusterId(href) // TODO: both
			a.NumberOfCitations = parseNumberOfCitations(text)
		}
		if strings.HasPrefix(href, "/scholar?cluster") {
			a.NumberOfVersions = parseNumberOfVersions(text) // TODO: fix
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
	a.PDFSource = parsePDFSource(sideBarA.Text())
}

/*
func (a *Article) crawlAndParseBibTeX() {
	popURL, err := CitePopUpQuery(a.InfoId)
	if err != nil {
	}

	popDoc, err := goquery.NewDocument(popURL)
	if err != nil {
	}

	bibURL, _ := popDoc.Find("#gs_citi > a:first-child").Attr("href")
	bibDoc, err := goquery.NewDocument(SCHOLAR_URL + bibURL)
	if err != nil {
	}
	a.Bibtex = bibDoc.Text()
}
*/

func (a *Article) String() string {
	title := fmt.Sprintf("title: %v\n", a.Title)
	year := fmt.Sprintf("year: %v\n", a.Year)
	url := fmt.Sprintf("url: %v\n", a.URL)
	cluster_id := fmt.Sprintf("cluster_id: %v\n", a.ClusterId)
	num_citations := fmt.Sprintf("# of citations: %v\n", a.NumberOfCitations)
	num_versions := fmt.Sprintf("$ of versions: %v\n", a.NumberOfVersions)
	info_id := fmt.Sprintf("info id: %v\n", a.InfoId)
	pdf_link := fmt.Sprintf("pdf link: %v\n", a.PDFLink)
	pdf_source := fmt.Sprintf("pdfSource: %v", a.PDFSource)
	ret := title + year + url + cluster_id + num_citations + num_versions + info_id + pdf_link + pdf_source

	return ret
}

func (a *Article) Json() string {
	bytes, err := json.Marshal(a)
	if err != nil {
		log.WithFields(log.Fields{"a": a, "err":err}).Error("Json encoding failed")
	}
	return string(bytes)
}

func (a *Article) IsValid() bool {
	// Avlid author-contamination. See #29 for details.
	title_validation := strings.HasPrefix(a.Title, "User profiles for")
	url_validation := strings.HasPrefix(a.URL, "/citations?view_op=search_authors")
	if title_validation && url_validation {
		return false
	}

	return true
}

func (a *Article) Same(b *Article) bool {
	title := a.Title == b.Title
	year := a.Year == b.Year
	url := a.hasSameURL(b)
	cluster_id := a.ClusterId == b.ClusterId
	number_of_citations := a.NumberOfCitations == b.NumberOfCitations // TODO: fix
	number_of_versions := a.NumberOfVersions == b.NumberOfVersions // TODO: fix
	info_id := a.InfoId == b.InfoId
	pdf_link := a.PDFLink == b.PDFLink
	pdf_source := a.PDFSource == b.PDFSource

	return title && year && url && cluster_id && number_of_citations && number_of_versions && info_id && pdf_link && pdf_source
}

func (a *Article) showDifference(b *Article) {
	if a.Title != b.Title {
		fmt.Println(a.Title)
		fmt.Println(b.Title)
	}
	if a.Year != b.Year {
		fmt.Println(a.Year)
		fmt.Println(b.Year)
	}
	if !a.hasSameURL(b) {
		if strings.HasPrefix(a.URL, "https://books.google.co.jp/") {
			trimParameter(trimParameter(a.URL, "sig"), "ots")
			trimParameter(trimParameter(b.URL, "sig"), "ots")
		} else {
			fmt.Println(a.URL)
			fmt.Println(b.URL)
		}
	}
	if a.ClusterId != b.ClusterId {
		fmt.Println(a.ClusterId)
		fmt.Println(b.ClusterId)
	}
	if a.NumberOfCitations != b.NumberOfCitations { // TODO: fix
		fmt.Println(a.NumberOfCitations)
		fmt.Println(b.NumberOfCitations)
	}
	if a.NumberOfVersions != b.NumberOfVersions { // TODO: fix
		fmt.Println(a.NumberOfVersions)
		fmt.Println(b.NumberOfVersions)
	}
	if a.InfoId != b.InfoId {
		fmt.Println(a.InfoId)
		fmt.Println(b.InfoId)
	}
	if a.PDFLink != b.PDFLink {
		fmt.Println(a.PDFLink)
		fmt.Println(b.PDFLink)
	}
	if a.PDFSource != b.PDFSource {
		fmt.Println(a.PDFSource)
		fmt.Println(b.PDFSource)
	}
}

func (a *Article) hasSameURL(b *Article) bool {
	if strings.HasPrefix(a.URL, "https://books.google.co.jp/") {
		return trimParameter(trimParameter(a.URL, "sig"), "ots") == trimParameter(trimParameter(b.URL, "sig"), "ots")
	}
	return a.URL == b.URL
}