package controller

import (
	"github.com/historical-rate/internal/app/adapter"
	"github.com/historical-rate/internal/app/application/usecase"
	"net/http"
)

func (s *Server) GetLatestHistoricalRate(w http.ResponseWriter, req *http.Request, ps adapter.Params) {
	rateUseCase := usecase.RateUseCase{RateRepository: s.RateRepository}
	rate, err := rateUseCase.GetLatestHistoricalRate()
	if err != nil {
		HandleBadRequest(w, err)
		return
	}
	response := make(map[string]interface{})
	response["rates"] = rate
	response["base"] = "EUR"
	HandleSuccessResponse(w, response)
}

func (s *Server) GetHistoricalRateByDate(w http.ResponseWriter, req *http.Request, ps adapter.Params) {
	date := ps.ByName("date")
	rateUseCase := usecase.RateUseCase{RateRepository: s.RateRepository}
	rate, err := rateUseCase.GetHistoricalRateByDate(date)
	if err != nil {
		HandleBadRequest(w, err)
		return
	}
	response := make(map[string]interface{})
	response["rates"] = rate
	response["base"] = "EUR"
	HandleSuccessResponse(w, response)
}

func (s *Server) GetHistoricalAnalyzeReport(w http.ResponseWriter, req *http.Request, ps adapter.Params) {
	rateUseCase := usecase.RateUseCase{RateRepository: s.RateRepository}
	rate, err := rateUseCase.GetHistoricalAnalyzes()
	if err != nil {
		HandleBadRequest(w, err)
		return
	}
	response := make(map[string]interface{})
	response["rates_analyze"] = rate
	response["base"] = "EUR"
	HandleSuccessResponse(w, response)
}
