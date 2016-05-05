package main

import (
	"testing"
	"fmt"
)

func TestParsePDFSource(t *testing.T) {
	s := "[PDF] from arxiv.orgarxiv.org [PDF]"
	expected := "arxiv.org"

	PDFLink := parsePDFSource(s)
	if PDFLink != expected {
		t.Error(fmt.Sprintf("\nExpected: %v\n  Actual: %v", expected, PDFLink))
	}
}

func TestTrimParameter(t *testing.T) {
	url1 := "https://books.google.co.jp/books?hl=ja&lr=&id=Tbn1l9P1220C&oi=fnd&pg=PA153&dq=deep+learning+%22y+bengio%22&ots=V3q8Fins1Z&sig=Z0htVnCeqaHiY7YLVRmqLJsZiBw#v=onepage&q&f=false"
	url2 := "https://books.google.co.jp/books?hl=en&lr=&id=y8ORL3DWt4sC&oi=fnd&pg=PR13&ots=bKyS8zP5Iz&sig=dC5YzrzUAz8kjnEx392vrjb6cr0"
	url3 := "https://books.google.co.jp/books?hl=ja&lr=&id=Tbn1l9P1220C&oi=fnd&pg=PA153&dq=deep+learning+%22y+bengio%22&ots=V3q8Fins1Z&sig=Z0htVnCeqaHiY7YLVRmqLJsZiBw&q&f=false"
	url4 := "https://books.google.co.jp/books?hl=en&lr=&id=y8ORL3DWt4sC&oi=fnd&pg=PR13&ots=bKyS9wNaHC&sig=wi01aFoEeNwUeehXa3OpNVjvLI0"
	url5 := "https://books.google.co.jp/books?hl=en&lr=&id=y8ORL3DWt4sC&oi=fnd&pg=PR13&ots=bKyS8zP5Iz&sig=dC5YzrzUAz8kjnEx392vrjb6cr0"
	url6 := "https://books.google.co.jp/books?hl=en&lr=&id=y8ORL3DWt4sC&oi=fnd&pg=PR13&ots=bKyS9xV3Fy&sig=bjLpIzuFfjnt_LIDwFx1S-1mg7w"
	url7 := "https://books.google.co.jp/books?hl=en&lr=&id=y8ORL3DWt4sC&oi=fnd&pg=PR13&ots=bKyS8zP5Iz&sig=dC5YzrzUAz8kjnEx392vrjb6cr0"

	expected1 := "https://books.google.co.jp/books?hl=ja&lr=&id=Tbn1l9P1220C&oi=fnd&pg=PA153&dq=deep+learning+%22y+bengio%22&ots=V3q8Fins1Z#v=onepage&q&f=false"
	expected2 := "https://books.google.co.jp/books?hl=en&lr=&id=y8ORL3DWt4sC&oi=fnd&pg=PR13&ots=bKyS8zP5Iz"
	expected3 := "https://books.google.co.jp/books?hl=ja&lr=&id=Tbn1l9P1220C&oi=fnd&pg=PA153&dq=deep+learning+%22y+bengio%22&ots=V3q8Fins1Z&q&f=false"

	if trimmed1 := trimParameter(url1, "sig"); trimmed1 != expected1 {
		t.Error(fmt.Sprintf("\nExpected: %v\n  Actual: %v", expected1, trimmed1))
	}
	if trimmed2 := trimParameter(url2, "sig"); trimmed2 != expected2 {
		t.Error(fmt.Sprintf("\nExpected: %v\n  Actual: %v", expected2, trimmed2))
	}
	if trimmed3 := trimParameter(url3, "sig"); trimmed3 != expected3 {
		t.Error(fmt.Sprintf("\nExpected: %v\n  Actual: %v", expected3, trimmed3))
	}

	trimmed4 := trimParameter(trimParameter(url4, "ots"), "sig")
	trimmed5 := trimParameter(trimParameter(url5, "ots"), "sig")
	if trimmed4 != trimmed5 {
		t.Error(fmt.Sprintf("\nurl4: %v\nurl5: %v", trimmed4, trimmed5))
	}
	trimmed6 := trimParameter(trimParameter(url6, "ots"), "sig")
	trimmed7 := trimParameter(trimParameter(url7, "ots"), "sig")
	if trimmed6 != trimmed7 {
		t.Error(fmt.Sprintf("\nurl6: %v\nurl7: %v", trimmed6, trimmed7))
	}
}