package goscholar

import (
	"fmt"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"strings"
	"strconv"
)

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
	Name string
	Url string
}

type Link struct {
	Name string
	Url  string
	Format string
}

// NewArticle creates an Article in which all entry is blank
func NewArticle() *Article {
	a := &Article{}
	title := &Title{}
	link := &Link{}

	a.Title = title
	a.Link = link

	return a
}

func (a *Article) String() string {
	ret := "-----------------------------------------------------------------------------\n"
	ret += "[Title]\n"
	ret += fmt.Sprintf("  Name: %v\n", a.Title.Name)
	ret += fmt.Sprintf("  Url: %v\n", a.Title.Url)
	ret += fmt.Sprintf("[Year]\n  %v\n", a.Year)
	ret += fmt.Sprintf("[ClusterId]\n  %v\n", a.ClusterId)
	ret += fmt.Sprintf("[NumeCite]\n  %v\n", a.NumCite)
	ret += fmt.Sprintf("[NumVer]\n  %v\n", a.NumVer)
	ret += fmt.Sprintf("[InfoId]\n  %v\n", a.InfoId)
	ret += "[Link]\n"
	ret += fmt.Sprintf("  Name: %v\n", a.Link.Name)
	ret += fmt.Sprintf("  Url: %v\n", a.Link.Url)
	ret += fmt.Sprintf("  Format: %v", a.Link.Format)

	return ret
}

func (a *Article) Json() string {
	bytes, err := json.Marshal(a)
	if err != nil {
		log.WithFields(log.Fields{"a": a, "err":err}).Error("Json encoding failed")
	}
	return string(bytes)
}

func (a *Article) isValid() bool { // TODO: fix (return error w/ message)
	// avoid author-contamination. See #29 for details.
	title_validation := strings.HasPrefix(a.Title.Name, "User profiles for")
	url_validation := strings.HasPrefix(a.Title.Url, "/citations?view_op=search_authors")
	if title_validation && url_validation {
		return false
	}

	// detect wrong year
	yearInt, err := strconv.Atoi(a.Year)
	if err != nil {
		return false
	}
	if !(1800 <= yearInt && yearInt <= 2100) {
		return false
	}

	return true
}