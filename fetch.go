package goscholar

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strings"
)

// Fetch gets a Document from a given URL. For usage, see the example of Overview.
func Fetch(url string) (doc *goquery.Document, err error) {
	log.WithFields(log.Fields{"url": url}).Info("Fetch sends request")

	// set request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.WithFields(log.Fields{"url": url, "err": err}).Error("Failed to generate new request")
		return nil, err
	}
	req.Header.Set("User-Agent", USER_AGENT)

	// send request and get response
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.WithFields(log.Fields{"url": url, "err": err}).Error("Failed to get response")
		return nil, err
	}

	// generate new Document
	doc, err = goquery.NewDocumentFromResponse(res)
	log.WithFields(log.Fields{"doc.url": doc.Url}).Info("goquery.Document is generated")
	if err != nil {
		log.WithFields(log.Fields{"url": url, "err": err}).Error("Generating goquery.Documentation failed")
		return nil, err
	}

	// 1. check the "Please show you're not a robot" page. See #61
	// 2. check the "We're sorry..."
	if s := doc.Find("h1").First().Text(); strings.Contains(s, "robot") || strings.Contains(s, "sorry") {
		log.WithFields(log.Fields{"h1": s, "doc.Url": doc.Url}).Error("Robot check occurs")
		return nil, errors.New("Failed to fetch Document")
	}

	// check the "To continue, please type the characters below:". See #55
	if strings.Contains(doc.Url.String(), "sorry") {
		log.WithFields(log.Fields{"doc.Url": doc.Url}).Error("Request is rejected from Google")
		return nil, errors.New("Failed to fetch Document")
	}

	return doc, nil
}
