package goscholar

import (
	"fmt"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
)

type Article struct {
	Title             *title
	Year              string
	ClusterId         string
	NumCite           string
	NumVer            string
	InfoId            string
	Link              *link
}

type title struct {
	Name string
	Url string
}

type link struct {
	Name string
	Url  string
	Format string
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
	ret += "\n-----------------------------------------------------------------------------"

	return ret
}

func (a *Article) Json() string {
	bytes, err := json.Marshal(a)
	if err != nil {
		log.WithFields(log.Fields{"a": a, "err":err}).Error("Json encoding failed")
	}
	return string(bytes)
}