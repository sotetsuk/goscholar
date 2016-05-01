package main

import (
	"strings"
)

const (
	SCHOLAR_URL = "https://scholar.google.co.jp/"
	WHOLE_ARTICLE_SELECTOR   = ".gs_r"
	ARTICLE_TITLE_SELECTOR   = ".gs_rt > a"
	ARTICLE_HEADER_SELECTOR  = ".gs_a"
	ARTICLE_FOOTER_SELECTOR  = ".gs_fl"
	ARTICLE_SIDEBAR_SELECTOR = ".gs_md_wp > a"
)

func parseClusterId(url string) string {
	url = url[15:]
	ix := strings.Index(url, "&")
	url = url[:ix]
	return url
}

func parseNumberOfCitations(s string) string {
	return s[8:]
}

func parseNumberOfVersions(s string) string {
	// TODO: WRITE HERE
	return s
}

func parseInfoId(url string) string {
	url = url[19:]
	ix := strings.Index(url, ":")
	url = url[:ix]
	return url
}

