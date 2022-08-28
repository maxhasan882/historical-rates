package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Header map[string]string

type Api struct {
	Url            string
	Params         string
	OptionalParams string
	Method         string
	Body           io.Reader
	Headers        []Header
}

func (a Api) MakeRequest() ([]byte, int, error) {
	url := fmt.Sprintf("%v%v%v", a.Url, a.Params, a.OptionalParams)
	request, err := http.NewRequest(a.Method, url, a.Body)
	for key := range a.Headers {
		for item, value := range a.Headers[key] {
			request.Header.Set(item, value)
		}
	}
	if err != nil {
		return []byte(""), 500, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return []byte(""), http.StatusInternalServerError, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(response.Body)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte(""), response.StatusCode, err
	}
	return body, response.StatusCode, nil
}
