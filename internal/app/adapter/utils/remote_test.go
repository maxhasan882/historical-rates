package utils

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestMakeRequestRequestError(t *testing.T) {
	api := Api{
		Url:            "error",
		Params:         "",
		OptionalParams: "",
		Method:         "",
		Body:           nil,
		Headers:        nil,
	}
	_, _, err := api.MakeRequest()
	assert.NotNil(t, err)
}

func TestMakeRequestClientError(t *testing.T) {
	api := Api{
		Url:            "https://httpbin.or/",
		Params:         "",
		OptionalParams: "",
		Method:         "",
		Body:           nil,
		Headers:        nil,
	}
	_, code, err := api.MakeRequest()
	assert.NotNil(t, err)
	assert.Equal(t, code, http.StatusInternalServerError)
}

func TestMakeRequestClientSuccess(t *testing.T) {
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	api := Api{
		Url:            "https://www.google.com/",
		Params:         "",
		OptionalParams: "",
		Method:         "",
		Body:           nil,
		Headers:        []Header{header},
	}
	_, code, err := api.MakeRequest()
	assert.Equal(t, err, nil)
	assert.Equal(t, code, http.StatusOK)
}
