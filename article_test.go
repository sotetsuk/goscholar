package goscholar

import (
	"testing"
)

var a *Article

func init() {
	a = &Article{
		Title: &title{
			Name: "Deep learning via Hessian-free optimization",
			Url:  "http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf"},
		Year:      "2010",
		ClusterId: "15502119379559163003",
		NumCite:   "260",
		NumVer:    "9",
		InfoId:    "e6RSJHGXItcJ",
		Link: &link{
			Name:   "wustl.edu",
			Url:    "http://machinelearning.wustl.edu/mlpapers/paper_files/icml2010_Martens10.pdf",
			Format: "PDF",
		},
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
  Format: PDF
-----------------------------------------------------------------------------`

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
