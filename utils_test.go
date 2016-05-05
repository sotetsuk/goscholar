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