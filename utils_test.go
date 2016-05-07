package goscholar

import (
	"testing"
)

func TestSame(t *testing.T) {
	a := NewArticle()
	b := NewArticle()

	if !same(a, b) {
		t.Error("Two Articles should be same")
	}

	b.Title.Name = "The title of article b"

	if same(a, b) {
		t.Error("Two Articles should be differetn")
	}
}

func ExampleShowDifference() {
	a := NewArticle()
	b := NewArticle()

	a.Title.Name = "Name of A"
	b.Title.Name = "Name of B"

	if !same(a, b) {
		showDifference(a, b)
	}
	// Output:
	// Title.Name is different
	// Name of A
	// Name of B
}
