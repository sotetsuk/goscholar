package goscholar

type Article struct {
	Title             *Title
	Year              string
	ClusterId         string
	NumCite           string
	NumVer            string
	InfoId            string
	Link              *Link
}

type Title struct {
	name string
	url string
}

type Link struct {
	name string
	url  string
	format string
}

func (a *Article) String() string {
	return ""
}

func (a *Article) Json() string {
	return ""
}