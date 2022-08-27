package controller

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/historical-rate/internal/app/adapter"
	"github.com/historical-rate/internal/app/application/usecase"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLatestHistoricalRate(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := usecase.NewMockIHistoricalRate(ctrl)
	m.EXPECT().GetLatest().Return(nil, nil)

	server := Server{
		DB:             nil,
		RateRepository: m,
	}
	server.RateRepository = m
	var response interface{}

	req := httptest.NewRequest(http.MethodGet, "/rates/latest", nil)
	res := httptest.NewRecorder()

	server.GetLatestHistoricalRate(res, req, nil)
	err := json.Unmarshal([]byte(res.Body.String()), &response)

	if err != nil {
		t.Error("Parse JSON Data Error")
	}
	if want, got := http.StatusOK, res.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}

	assert.Equal(t, res.Code, 200)
}

func TestHistoricalRateByDate(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := usecase.NewMockIHistoricalRate(ctrl)
	date := "2022-01-01"
	m.EXPECT().GetByDate(date).Return(nil, nil)

	server := Server{
		DB:             nil,
		RateRepository: m,
	}
	server.RateRepository = m
	var response interface{}

	req := httptest.NewRequest(http.MethodGet, "/rates/"+date, nil)

	res := httptest.NewRecorder()

	server.GetHistoricalRateByDate(res, req, []adapter.Param{{
		Key:   "date",
		Value: date,
	}})
	err := json.Unmarshal([]byte(res.Body.String()), &response)

	if err != nil {
		t.Error("Parse JSON Data Error")
	}
	assert.Equal(t, res.Code, 200)
}

func TestHistoricalAnalyzeReport(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := usecase.NewMockIHistoricalRate(ctrl)
	m.EXPECT().GetAnalyze().Return(nil, nil)

	server := Server{
		DB:             nil,
		RateRepository: m,
	}
	server.RateRepository = m
	var response interface{}

	req := httptest.NewRequest(http.MethodGet, "/rates/analyze", nil)
	res := httptest.NewRecorder()

	server.GetHistoricalAnalyzeReport(res, req, nil)
	err := json.Unmarshal([]byte(res.Body.String()), &response)

	if err != nil {
		t.Error("Parse JSON Data Error")
	}
	if want, got := http.StatusOK, res.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}

	assert.Equal(t, res.Code, 200)
}
