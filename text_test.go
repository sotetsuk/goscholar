package goscholar

import (
	"testing"
	"fmt"
)

func TestParseYearText(t *testing.T) {
	s := "Y LeCun, Y Bengio, G Hinton - Nature, 2015 - nature.com"
	expected := "2015"

	if actual := parseYearText(s); actual != expected {
		t.Error(TestErr{expected:expected, actual:actual})
	}
}

func TestParseClusterIdText(t *testing.T) {
	url := "/scholar?cites=5362332738201102290&as_sdt=2005&sciodt=0,5&hl=en"
	expected := "5362332738201102290"

	if actual := parseClusterIdText(url); actual != expected {
		t.Error(TestErr{expected:expected, actual:actual})
	}
}

func TestParseNumCiteText(t *testing.T) {
	s := "Cited by 390"
	expected := "390"

	if actual := parseNumCiteText(s); actual != expected {
		t.Error(TestErr{expected:expected, actual:actual})
	}
}

func TestParseNumVerText(t *testing.T) {
	s := "All 7 versions"
	expected := "7"

	if actual := parseNumVerText(s); actual != expected {
		t.Error(TestErr{expected:expected, actual:actual})
	}
}

func TestParseInfoIdText(t *testing.T) {
	url := "/scholar?q=related:0qfs6zbVakoJ:scholar.google.com/&hl=en&as_sdt=0,5"
	expected := "0qfs6zbVakoJ"

	if actual := parseInfoIdText(url); actual != expected {
		t.Error(TestErr{expected:expected, actual:actual})
	}
}

func TestParseLinkText(t *testing.T) {
	s := "psu.edu [PDF]"
	expectedName := "psu.edu"
	expectedFormat := "PDF"

	actualName, actualFormat := parseLinkText(s)

	if actualName != expectedName {
		t.Error(TestErr{expected:expectedName, actual:actualName})
	}

	if actualFormat != expectedFormat {
		t.Error(TestErr{expected:expectedFormat, actual:actualFormat})
	}
}

func TestEnclosedInDoubleQuotation(t *testing.T) {
	s1 := "y bengio"
	s2 := "\"y bengio\""

	if enclosedInDoubleQuotation(s1) {
		t.Error("%v is not enclosed in double quotation")
	}

	if !enclosedInDoubleQuotation(s2) {
		t.Error("%v is enclosed in double quotation")
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