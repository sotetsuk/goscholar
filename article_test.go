package goscholar

import (
	"testing"
	"fmt"
	"errors"
)

var a *Article

func init() {
	a = &Article{
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
	a := NewArticle()

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

func TestString(t *testing.T) {
	expected := `-----------------------------------------------------------------------------
[Title]
  Name: Deep learning via Hessian-free optimization
  Url: http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf
[Year]
  2010
[ClusterId]
  15502119379559163003
[NumeCite]
  260
[NumVer]
  9
[InfoId]
  e6RSJHGXItcJ
[Link]
  Name: wustl.edu
  Url: http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf
  Format: PDF`

	if actual := a.String(); actual != expected {
		t.Error(TestErr{expected:expected, actual:actual})
	}
}

func TestJson(t *testing.T) {
	expected := `{"Title":{"Name":"Deep learning via Hessian-free optimization","Url":"http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf"},"Year":"2010","ClusterId":"15502119379559163003","NumCite":"260","NumVer":"9","InfoId":"e6RSJHGXItcJ","Link":{"Name":"wustl.edu","Url":"http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf","Format":"PDF"}}`

	if actual := a.Json(); actual != expected {
		t.Error(TestErr{expected:expected, actual:actual})
	}
}

func TestIsValid(t *testing.T) {
	// test case 1
	aInvalid := NewArticle()
	aInvalid.Title.Name = "User profiles for author:\"y bengio\""
	aInvalid.Title.Url = "/citations?view_op=search_authors&mauthors=author:%22y+bengio%22&hl=en&oi=ao"

	if aInvalid.isValid() {
		t.Error(fmt.Sprintf("\n%v \nThis title should be invalid", aInvalid.Title.Name))
	}

	aInvalid = NewArticle()
	a.Year = "1206"
	if aInvalid.isValid() {
		t.Error(fmt.Sprintf("\n%v \nThis year should be invalid", aInvalid.Year))
	}
}