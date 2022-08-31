package controller

import (
	"encoding/json"
	"net/http"
)

func HandleBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	_, err = w.Write([]byte(err.Error()))
}

func HandleSuccessResponse(w http.ResponseWriter, response map[string]interface{}) {
	rateJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
	}
	_, err = w.Write(rateJson)
}
