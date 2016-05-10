package goscholar

// URL is used by Query in order to issue a url to which a function Fetch sends request
const (
	scholar_url    = "https://scholar.google.co.jp/"
	search_url     = scholar_url + "scholar?hl=en&q=%v&as_ylo=%v&as_yhi=%v&num=%v&start=%v"
	find_url       = scholar_url + "scholar?hl=en&cluster=%v&num=1"
	cite_url       = scholar_url + "scholar?hl=en&cites=%v&as_ylo=%v&as_yhi=%v&num=%v&start=%v"
	cite_popup_url = scholar_url + "scholar?q=info:%s:scholar.google.com/&output=cite&scirp=0&hl=en"
)

// SELECTOR is used for parsing a fetched Document
const (
	whole_article_selector      = ".gs_r"
	article_h3_selector         = "h3.gs_rt > a"
	article_green_line_selector = ".gs_a"
	article_bottom_selector     = ".gs_fl"
	article_sidebar_selector    = ".gs_md_wp > a"
	sidebar_text_selector       = ".gs_ggsS"
)
