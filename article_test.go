package goscholar

import (
	"testing"
	"fmt"
	"errors"
)

var article *Article

func init() {
	article = &Article{
		Title: &Title{
			Name: "Deep learning via Hessian-free optimization",
			Url:  "http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf"},
		Year:      "2010",
		ClusterId: "15502119379559163003",
		NumCite:   "260",
		NumVer:    "9",
		InfoId:    "e6RSJHGXItcJ",
		Link: &Link{
			Name:   "wustl.edu",
			Url:    "http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf",
			Format: "PDF",
		},
	}
}

func TestNewArticle(t *testing.T) {
	a := newArticle()

	title := a.Title
	year := a.Year
	clusterId := a.ClusterId
	numCite := a.NumCite
	numVer := a.NumVer
	infoId := a.InfoId
	link := a.Link

	var err error
	checkBlank := func(s string) error {
		if err != nil {
			return err
		}

		if s != "" {
			err = errors.New(fmt.Sprintf("%v is not blank", s))
		}

		return err
	}

	checkBlank(title.Name)
	checkBlank(title.Url)
	checkBlank(year)
	checkBlank(clusterId)
	checkBlank(numCite)
	checkBlank(numVer)
	checkBlank(infoId)
	checkBlank(link.Name)
	checkBlank(link.Url)
	checkBlank(link.Format)

	if err != nil {
		t.Error(err)
	}
}

func ExampleString() {
	fmt.Println(article)
	// Output:
	/*[Title]
  Name: Deep learning via Hessian-free optimization
  Url: http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf
[Year]
  2010
[ClusterId]
  15502119379559163003
[NumCite]
  260
[NumVer]
  9
[InfoId]
  e6RSJHGXItcJ
[Link]
  Name: wustl.edu
  Url: http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf
  Format: PDF*/
}

func ExampleJson() {
	fmt.Println(article.Json())
	// Output:
	// {"title":{"name":"Deep learning via Hessian-free optimization","url":"http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf"},"year":"2010","cluster_id":"15502119379559163003","num_cite":"260","num_ver":"9","info_id":"e6RSJHGXItcJ","link":{"name":"wustl.edu","url":"http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf","format":"PDF"}}
}

func TestIsValid(t *testing.T) {
	// test case 1
	aInvalid := newArticle()
	aInvalid.Title.Name = "User profiles for author:\"y bengio\""
	aInvalid.Title.Url = "/citations?view_op=search_authors&mauthors=author:%22y+bengio%22&hl=en&oi=ao"

	if aInvalid.isValid() {
		t.Error(fmt.Sprintf("\n%v \nThis title should be invalid", aInvalid.Title.Name))
	}

	aInvalid = newArticle()
	aInvalid.Year = "1206"
	if aInvalid.isValid() {
		t.Error(fmt.Sprintf("\n%v \nThis year should be invalid", aInvalid.Year))
	}
}