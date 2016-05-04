package main

import (
	"strings"
	"regexp"
)


func parseYear(s string) string {
	re, _ := regexp.Compile("\\d{4}")
	return strings.TrimSpace(string(re.Find([]byte(s))))
}

func parseClusterId(url string) string {
	url = url[15:]
	ix := strings.Index(url, "&")
	url = url[:ix]
	return strings.TrimSpace(url)
}

func parseNumberOfCitations(s string) string {
	return strings.TrimSpace(s[8:])
}

func parseNumberOfVersions(s string) string {
	return strings.TrimSpace(s[strings.Index(s, " "):strings.LastIndex(s, " ")])
}

func parseInfoId(url string) string {
	url = url[19:]
	ix := strings.Index(url, ":")
	url = url[:ix]
	return strings.TrimSpace(url)
}

func parsePDFSource(s string) string {
	// e.g., "[PDF] from arxiv.orgarxiv.org [PDF]"", => "PDFSource": "arxiv.org"
	prefix := "[PDF] from "
	suffix := " [PDF]"
	if strings.HasPrefix(s, prefix) && strings.HasSuffix(s, suffix) {
		s = strings.TrimSpace(s[len(prefix):len(s) - len(suffix)])
		return s[:len(s) / 2]
	}
	return ""
}

func StartAndEndWithDoubleQuotation(s string) bool {
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		return true
	} else {
		return false
	}
}