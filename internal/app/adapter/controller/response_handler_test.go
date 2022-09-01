package controller

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleBadRequest(t *testing.T) {
	res := httptest.NewRecorder()
	err := "test error"
	HandleBadRequest(res, errors.New(err))
	assert.Equal(t, res.Body.String(), err)
	assert.Equal(t, res.Code, http.StatusBadRequest)
}

func TestHandleSuccessResponseValidData(t *testing.T) {
	res := httptest.NewRecorder()
	body := make(map[string]interface{})
	body["message"] = "test message"
	HandleSuccessResponse(res, body)
	response := make(map[string]interface{})
	err := json.Unmarshal([]byte(res.Body.String()), &response)
	if err != nil {
		t.Error("Parse JSON Data Error")
	}
	assert.Equal(t, body, response)
	assert.Equal(t, res.Code, http.StatusOK)
}

func TestHandleSuccessResponseInvalidData(t *testing.T) {
	res := httptest.NewRecorder()
	ss := make(map[string]interface{})
	ss["hello"] = make(chan int)
	HandleSuccessResponse(res, ss)
	assert.Equal(t, res.Code, http.StatusBadRequest)
}
