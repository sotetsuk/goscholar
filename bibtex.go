package goscholar

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func getBibTeX(url string) (bibtex string, err error) {
	// TODO: add logging
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{Name: "GSP", Value: "CF=4"})

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	// e.g., 403 Forbidden occurs when we remove "GSP=CF=4"
	if res.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("Bad status code: %v", res.StatusCode))
	}

	// convert to string
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(res.Body)
	body := bufbody.String()

	// trim \n and abundant spaces
	body = strings.Replace(body, "\n", "", -1)
	body = strings.Replace(body, "  ", " ", -1)

	return body, nil
}
