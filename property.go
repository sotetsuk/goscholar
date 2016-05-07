package goscholar

// url
const (
	SCHOLAR_URL = "https://scholar.google.co.jp/"
	SEARCH_URL     = SCHOLAR_URL + "scholar?hl=en&q=%v&as_ylo=%v&as_yhi=%v&num=%v&start=%v"
	FIND_URL       = SCHOLAR_URL + "scholar?hl=en&cluster=%v&num=1"
	CITE_URL       = SCHOLAR_URL + "scholar?hl=en&cites=%v&as_ylo=%v&as_yhi=%v&num=%v&start=%v"
	CITE_POPUP_URL = SCHOLAR_URL + "scholar?q=info:%s:scholar.google.com/&output=cite&scirp=0&hl=en"
)

// selector
const (
	WHOLE_ARTICLE_SELECTOR = ".gs_r"
	ARTICLE_TITLE_SELECTOR   = "h3.gs_rt > a"
	ARTICLE_HEADER_SELECTOR  = ".gs_a"
	ARTICLE_FOOTER_SELECTOR  = ".gs_fl"
	ARTICLE_SIDEBAR_SELECTOR = ".gs_md_wp > a"
)