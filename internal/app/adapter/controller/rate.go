package controller

import (
	"encoding/json"
	"github.com/historical-rate/internal/app/application/usecase"
	"net/http"
)

func (s *Server) GetLatestHistoricalRate(w http.ResponseWriter, req *http.Request) {
	rateUseCase := usecase.RateUseCase{RateRepository: s.RateRepository}
	rate, err := rateUseCase.GetLatestHistoricalRate()
	response := make(map[string]interface{})
	response["rates"] = rate
	response["base"] = "EUR"
	rateJson, err := json.Marshal(response)
	_, err = w.Write(rateJson)
	if err != nil {
		return
	}
}
