package utils

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Header map[string]string

// Api request attribute schema
type Api struct {
	Url            string
	Params         string
	OptionalParams string
	Method         string
	Body           io.Reader
	Headers        []Header
}

// MakeRequest request to external source and returns payload data, response status and error if any.
func (a Api) MakeRequest() ([]byte, int, error) {
	url := fmt.Sprintf("%v%v%v", a.Url, a.Params, a.OptionalParams)
	request, err := http.NewRequest(a.Method, url, a.Body)
	if err != nil {
		return []byte(""), 500, err
	}
	for key := range a.Headers {
		for item, value := range a.Headers[key] {
			request.Header.Set(item, value)
		}
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // for overlooking certificate signed by unknown authority x509 issue
	}
	client := &http.Client{Transport: tr}
	response, err := client.Do(request)
	if err != nil {
		return []byte(""), http.StatusInternalServerError, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return body, response.StatusCode, err
}
