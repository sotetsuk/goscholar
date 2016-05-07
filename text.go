package goscholar

import (
	"regexp"
	"strings"
	"fmt"
)

func parseYearText(s string) (year string) {
	re, _ := regexp.Compile("\\d{4}")

	return strings.TrimSpace(string(re.Find([]byte(s))))

}

func parseClusterIdText(url string) (clusterId string) {
	clusterId = strings.TrimPrefix(url, "/scholar?cites=")
	clusterId = clusterId[:strings.Index(clusterId, "&")]

	return strings.TrimSpace(clusterId)
}

func parseNumCiteText(s string) (numCite string) {
	s = strings.TrimPrefix(s, "Cited by")

	return strings.TrimSpace(s)
}

func parseNumVerText(s string) (numVer string) {
	s = strings.TrimPrefix(s, "All")
	s = strings.TrimSuffix(s, "versions")

	return strings.TrimSpace(s)
}

func parseInfoIdText(url string) (infoId string) {
	infoId = strings.TrimPrefix(url, "/scholar?q=related:")
	infoId = infoId[:strings.Index(infoId, ":scholar.google.com")]

	return strings.TrimSpace(infoId)
}

func parseLinkText(s string) (name, format string) {
	ix := strings.Index(s, " ")
	name = strings.TrimSpace(s[:ix])
	format = strings.TrimSpace(s[ix:])

	return name, format[1 : len(format)-1]
}

// checkDoubleQuotation return true if s starts and ends with double quotation
// E.g. checkDoubleQuotation("\"y bengio\"") => true
func enclosedInDoubleQuotation(s string) bool {
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		return true
	} else {
		return false
	}
}

func trimParameter(url string, param string) string {
	rep := regexp.MustCompile(fmt.Sprintf(`&%v=[A-Za-z0-9_-]*`, param))
	return rep.ReplaceAllString(url, "")
}
