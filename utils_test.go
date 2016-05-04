package main

import (
	"testing"
	"fmt"
)

func TestParsePDFLink(t *testing.T) {
	s := "[PDF] from arxiv.orgarxiv.org [PDF]"
	expected := "arxiv.org"

	PDFLink := parsePDFLink(s)
	if PDFLink != expected {
		t.Error(fmt.Sprintf("TestParsePdfLink failed:\n  Expected: %v\n    Actual: %v", expected, PDFLink))
	}
	fmt.Println(fmt.Sprintf("PDFLink: %v", PDFLink))
}