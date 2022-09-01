package controller

import (
	"encoding/json"
	"net/http"
)

// HandleBadRequest handles response for unsuccessful request
func HandleBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	_, err = w.Write([]byte(err.Error()))
}

// HandleSuccessResponse handles response for successful request
func HandleSuccessResponse(w http.ResponseWriter, response map[string]interface{}) {
	rateJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
	} else {
		_, err = w.Write(rateJson)
	}

}
