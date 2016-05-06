package goscholar
import "strings"

func parseYearText(s string) string {
	return ""
}

func parseClusterIdText(url string) string {
	return ""
}

func parseNumberOfCitationsText(s string) string {
	return ""
}

func parseNumberOfVersionsText(s string) string {
	return ""
}

func parseInfoIdText(url string) string {
	return ""
}

func parsePDFSourceText(s string) string {
	return ""
}

// checkDoubleQuotation return true if s starts and ends with double quotation
// E.g. checkDoubleQuotation("\"y bengio\"") => true
func checkDoubleQuotation(s string) bool {
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		return true
	} else {
		return false
	}
}

func trimParameter(url string, param string) string {
	return ""
}