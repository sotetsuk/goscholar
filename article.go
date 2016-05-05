package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
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
	url := a.URL == b.URL
	cluster_id := a.ClusterId == b.ClusterId
	number_of_citations := a.NumberOfCitations == b.NumberOfCitations // TODO:
	number_of_versions := a.NumberOfVersions == b.NumberOfVersions // TODO:
	info_id := a.InfoId == b.InfoId
	pdf_link := a.PDFLink == b.PDFLink
	pdf_source := a.PDFSource == b.PDFSource

	return title && year && url && cluster_id && number_of_citations && number_of_versions && info_id && pdf_link && pdf_source
}