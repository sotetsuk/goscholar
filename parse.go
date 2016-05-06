package goscholar

import (
	"github.com/PuerkitoBio/goquery"
)

func ParseDocument(ch chan *Article, s *goquery.Document) {}

func ParseSelection(s *goquery.Selection) (a *Article, err error) {
	return nil, nil
}

func parseH3(s *goquery.Selection) (title *title){
	return nil
}

func parseGreenLine(s *goquery.Selection) (year string) {
	return ""
}

func parseBottom(s *goquery.Selection) (cluster_id, n_cite, n_ver, info_id string) {
	return "", "", "", ""
}

func parseSideBar(s *goquery.Selection) (link *link) {
	return nil
}