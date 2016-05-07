package goscholar

import (
	"testing"
)

func TestParseYearText(t *testing.T) {
	// test case 1
	s := "Y LeCun, Y Bengio, G Hinton - Nature, 2015 - nature.com"
	expected := "2015"

	if actual := parseYearText(s); actual != expected {
		t.Error(testErr{expected:expected, actual:actual})
	}

	// test case 1
	s = "Y Bengio, AC Courville, P Vincent - CoRR, abs/1206.5538, 2012"
	expected = "2012"
	if actual := parseYearText(s); actual != expected {
		t.Error(testErr{expected:expected, actual:actual})
	}
}

func TestParseClusterIdText(t *testing.T) {
	url := "/scholar?cites=5362332738201102290&as_sdt=2005&sciodt=0,5&hl=en"
	expected := "5362332738201102290"

	if actual := parseClusterIdText(url); actual != expected {
		t.Error(testErr{expected:expected, actual:actual})
	}
}

func TestParseNumCiteText(t *testing.T) {
	s := "Cited by 390"
	expected := "390"

	if actual := parseNumCiteText(s); actual != expected {
		t.Error(testErr{expected:expected, actual:actual})
	}
}

func TestParseNumVerText(t *testing.T) {
	s := "All 7 versions"
	expected := "7"

	if actual := parseNumVerText(s); actual != expected {
		t.Error(testErr{expected:expected, actual:actual})
	}
}

func TestParseInfoIdText(t *testing.T) {
	url := "/scholar?q=related:0qfs6zbVakoJ:scholar.google.com/&hl=en&as_sdt=0,5"
	expected := "0qfs6zbVakoJ"

	if actual := parseInfoIdText(url); actual != expected {
		t.Error(testErr{expected:expected, actual:actual})
	}
}

func TestParseLinkText(t *testing.T) {
	s := "psu.edu [PDF]"
	expectedName := "psu.edu"
	expectedFormat := "PDF"

	actualName, actualFormat := parseLinkText(s)

	if actualName != expectedName {
		t.Error(testErr{expected:expectedName, actual:actualName})
	}

	if actualFormat != expectedFormat {
		t.Error(testErr{expected:expectedFormat, actual:actualFormat})
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