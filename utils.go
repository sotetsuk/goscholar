package goscholar

import "fmt"

type TestErr struct {
	expected string
	actual string
}

func (e TestErr) Error() string {
	return fmt.Sprintf("\nExpected: %v\n  Actual: %v", e.expected, e.actual)
}

func same(a *Article, b *Article) (ok bool) {
	return true
}

func showDifference(a *Article, b *Article) {

}

func hasSameURL(a *Article, b *Article) (ok bool) {
	return true
}