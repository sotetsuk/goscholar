package goscholar

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	log "github.com/Sirupsen/logrus"
	"regexp"
	"strings"
)

// ParseDocument sends the pointers of parsed Articles to the given channel.
// The channel will be closed if there are no articles to be sent.
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
	doc.Find(whole_article_selector).Each(parse)
}

// ParseSelection returns a parsed Article.
// If the Article is not valid (e.g., Author profile), it returns error.
func ParseSelection(s *goquery.Selection) (a *Article, err error) {
	a = &Article{}

	a.Title = parseH3(s)
	a.Year = parseGreenLine(s)
	a.ClusterId, a.NumCite, a.NumVer, a.InfoId = parseBottom(s)
	a.Link = parseSideBar(s)

	if !a.isValid() {
		return nil, errors.New(fmt.Sprintf("\"%v\" is not a valid article", a.Title.Name))
	}

	return a, nil
}

// parseH3 an article title and its link
func parseH3(s *goquery.Selection) (title *Title) {
	title = &Title{}
	h3 := s.Find(article_h3_selector)
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
	year = parseYearText(s.Find(article_green_line_selector).Text())

	return year
}

// parseBottom parse the line under the abstract
func parseBottom(s *goquery.Selection) (clusterId, numCite, numVer, infoId string) {
	divFooter := s.Find(article_bottom_selector)
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

	sideBarA := s.Find(article_sidebar_selector)
	url, exists := sideBarA.Attr("href")
	if !exists {
		return link
	}

	link.Url = url
	link.Name, link.Format = parseLinkText(sideBarA.Find(sidebar_text_selector).Text())

	return link
}
