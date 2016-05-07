package goscholar

import (
	"github.com/PuerkitoBio/goquery"
	log "github.com/Sirupsen/logrus"
	"strings"
	"errors"
)

func Fetech(url string) (doc *goquery.Document, err error) {
	log.WithFields(log.Fields{"url": url}).Info("Fetch sends request")

	doc, err = goquery.NewDocument(url)
	log.WithFields(log.Fields{"doc.url": doc.Url}).Info("goquery.Document is generated")
	if err != nil {
		log.WithFields(log.Fields{"url": url, "err": err}).Error("Generating goquery.Documentation failed")
		return nil, err
	}

	// 1. check the "Please show you're not a robot" page. See #61
	// 2. check the "We're sorry..."
	if s := doc.Find("h1").First().Text(); strings.Contains(s, "robot") || strings.Contains(s, "sorry") {
		log.WithFields(log.Fields{"h1":s, "doc.Url": doc.Url}).Error("Robot check occurs")
		return nil, errors.New("Robot check occurs")
	}

	// check the "To continue, please type the characters below:". See #55
	if strings.Contains(doc.Url.String(), "sorry") {
		log.WithFields(log.Fields{"doc.Url": doc.Url}).Error("Request is rejected from Google")
		return nil, errors.New("Request is rejected from Google")
	}

	return doc, nil
}