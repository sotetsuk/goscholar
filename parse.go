package goscholar

import (
	"github.com/PuerkitoBio/goquery"
	log "github.com/Sirupsen/logrus"
	"regexp"
	"strings"
	"errors"
	"fmt"
)

func ParseDocument(ch chan *Article, doc *goquery.Document) {
	defer close(ch)

	parse := func(i int, s *goquery.Selection) {
		a, err := ParseSelection(s)
		if err != nil {
			log.Error(err)
			return
		}
		ch <- a
	}
	doc.Find(WHOLE_ARTICLE_SELECTOR).Each(parse)
}

// ParseSelection returns one parsed article.
// If article is not valid (e.g., Author profile), it returns error.
func ParseSelection(s *goquery.Selection) (a *Article, err error) {
	a = &Article{}

	a.Title = parseH3(s)
	a.Year = parseGreenLine(s)
	a.ClusterId, a.NumCite, a.NumVer, a.InfoId = parseBottom(s)
	a.Link = parseSideBar(s)

	if !a.isValid() {
		return nil, errors.New(fmt.Sprintf("%v is not valid article", a.Title.Name))
	}

	return a, nil
}

// parseH3 parse article title and its link
func parseH3(s *goquery.Selection) (title *Title){
	title = &Title{}
	h3 := s.Find(ARTICLE_TITLE_SELECTOR)
	url, exists := h3.Attr("href")

	if exists {
		title.Url = url
		title.Name = h3.Text()
	} else {
		name := s.Find("h3").Text()
		rep, _ := regexp.Compile("\\[[a-zA-Z0-9]*\\]\\[[a-zA-Z0-9]*\\]\\s")
		title.Name = rep.ReplaceAllString(name, "")
	}

	return title
}

// parseGreenLine parse article published year
func parseGreenLine(s *goquery.Selection) (year string) {
	year = parseYearText(s.Find(ARTICLE_HEADER_SELECTOR).Text())

	return year
}

// parseBottom parse the line under the abstract
func parseBottom(s *goquery.Selection) (clusterId, numCite, numVer, infoId string) {
	divFooter := s.Find(ARTICLE_FOOTER_SELECTOR)
	parseFooter := func(i int, s *goquery.Selection) {

		href, _ := s.Attr("href")
		text := s.Text()

		if strings.HasPrefix(href, "/scholar?cites") {
			clusterId = parseClusterIdText(href)
			numCite = parseNumCiteText(text)
		}
		if strings.HasPrefix(href, "/scholar?cluster") {
			numVer = parseNumVerText(text)
		}
		if strings.HasPrefix(href, "/scholar?q=related") {
			infoId = parseInfoIdText(href)
		}

	}
	divFooter.Find("a").Each(parseFooter)

	return clusterId, numCite, numVer, infoId
}

// parseSideBar parse the right side link
func parseSideBar(s *goquery.Selection) (link *Link) {
	link = &Link{}

	sideBarA := s.Find(ARTICLE_SIDEBAR_SELECTOR)
	url, exists := sideBarA.Attr("href")
	if !exists {
		return link
	}

	link.Url = url
	link.Name, link.Format = parseLinkText(sideBarA.Find(SIDEBAR_TEXT_SELECTOR).Text())

	return link
}