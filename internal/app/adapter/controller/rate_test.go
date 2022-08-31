package controller

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/historical-rate/internal/app/adapter"
	"github.com/historical-rate/internal/app/domain"
	"github.com/historical-rate/internal/app/domain/repository/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLatestHistoricalRate(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockIHistoricalRate(ctrl)
	timeNow := time.Now()
	var rates []domain.Rate
	rates = append(rates, domain.Rate{
		Id:       1,
		Date:     &timeNow,
		Currency: "BDT",
		Rate:     100,
	})

	expectedResult := make(map[string]interface{})
	for _, value := range rates {
		expectedResult[value.Currency] = value.Rate
	}
	actualResult := make(map[string]interface{})
	actualResult["base"] = "EUR"
	actualResult["rates"] = expectedResult

	m.EXPECT().GetLatest().Return(rates, nil)

	server := Server{
		DB:             nil,
		RateRepository: m,
	}
	server.RateRepository = m
	var response map[string]interface{}

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
	assert.Equal(t, response, actualResult)
	assert.Equal(t, res.Code, 200)
	mInv := mocks.NewMockIHistoricalRate(ctrl)
	serverInv := Server{
		DB:             nil,
		RateRepository: mInv,
	}
	server.RateRepository = mInv
	mInv.EXPECT().GetLatest().Return(rates, errors.New("new error"))
	req = httptest.NewRequest(http.MethodGet, "/rates/latest", nil)
	res = httptest.NewRecorder()
	serverInv.GetLatestHistoricalRate(res, req, nil)
	assert.Equal(t, res.Body.String(), "new error")
	assert.Equal(t, res.Code, 400)
}

func TestHistoricalRateByDate(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockIHistoricalRate(ctrl)
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
	m := mocks.NewMockIHistoricalRate(ctrl)
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
