package goscholar

import (
	"fmt"
)

type testErr struct {
	expected string
	actual   string
}

func (e testErr) Error() string {
	return fmt.Sprintf("\nExpected: %v\n  Actual: %v", e.expected, e.actual)
}

func same(a *Article, b *Article) (ok bool) {
	titleName := a.Title.Name == b.Title.Name
	titleUrl := a.Title.Url == b.Title.Url
	year := a.Year == b.Year
	cluster_id := a.ClusterId == b.ClusterId
	numCite := a.NumCite == b.NumCite // TODO: fix; check by range
	numVer := a.NumVer == b.NumVer    // TODO: fix; check by range
	info_id := a.InfoId == b.InfoId
	linkName := a.Link.Name == b.Link.Name
	linkUrl := a.Link.Url == b.Link.Url
	linkFormat := a.Link.Format == b.Link.Format

	return titleName && titleUrl && year && cluster_id && numCite && numVer && info_id && linkName && linkUrl && linkFormat
}

func showDifference(a *Article, b *Article) {
	if a.Title.Name != b.Title.Name {
		fmt.Println("Title.Name is different")
		fmt.Println(a.Title.Name)
		fmt.Println(b.Title.Name)
	}

	if a.Title.Url != b.Title.Url {
		fmt.Println("Title.Url is different")
		fmt.Println(a.Title.Url)
		fmt.Println(b.Title.Url)
	}

	if a.Year != b.Year {
		fmt.Println("Year is different")
		fmt.Println(a.Year)
		fmt.Println(b.Year)
	}
	if a.ClusterId != b.ClusterId {
		fmt.Println("ClusterId is different")
		fmt.Println(a.ClusterId)
		fmt.Println(b.ClusterId)
	}
	if a.NumCite != b.NumCite {
		fmt.Println("NumCite is different")
		fmt.Println(a.NumCite)
		fmt.Println(b.NumCite)
	}
	if a.NumVer != b.NumVer {
		fmt.Println("NumVer is different")
		fmt.Println(a.NumVer)
		fmt.Println(b.NumVer)
	}
	if a.InfoId != b.InfoId {
		fmt.Println("InfoId is different")
		fmt.Println(a.InfoId)
		fmt.Println(b.InfoId)
	}
	if a.Link.Name != b.Link.Name {
		fmt.Println("Link.Name is different")
		fmt.Println(a.Link.Name)
		fmt.Println(b.Link.Name)
	}
	if a.Link.Url != b.Link.Url {
		fmt.Println("Link.Url is different")
		fmt.Println(a.Link.Url)
		fmt.Println(b.Link.Url)
	}
	if a.Link.Format != b.Link.Format {
		fmt.Println("Title.Format is different")
		fmt.Println(a.Link.Format)
		fmt.Println(b.Link.Format)
	}
}
