package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/historical-rate/internal/app/domain"
	"github.com/historical-rate/internal/app/domain/repository/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetLatestHistoricalRate(t *testing.T) {
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
	expectedResult := make(map[string]float64)
	for _, value := range rates {
		expectedResult[value.Currency] = value.Rate
	}
	m.EXPECT().GetLatest().Return(rates, nil)
	res, err := RateUseCase{m}.GetLatestHistoricalRate()
	assert.Equal(t, res, expectedResult)
	assert.Equal(t, err, nil)
}

func TestGetLatestHistoricalRateError(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mocks.NewMockIHistoricalRate(ctrl)

	m.EXPECT().GetLatest().Return(nil, errors.New("test error"))
	_, err := RateUseCase{m}.GetLatestHistoricalRate()
	assert.NotNil(t, err)
}

func TestGetHistoricalRateByDate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockIHistoricalRate(ctrl)
	timeNow := time.Now()
	timeString := timeNow.Format("2006-01-03")
	var rates []domain.Rate
	rates = append(rates, domain.Rate{
		Id:       1,
		Date:     &timeNow,
		Currency: "BDT",
		Rate:     100,
	})

	result := make(map[string]float64)
	for _, value := range rates {
		result[value.Currency] = value.Rate
	}
	m.EXPECT().GetByDate(timeString).Return(rates, nil)
	res, err := RateUseCase{m}.GetHistoricalRateByDate(timeString)
	assert.Equal(t, res, result)
	assert.Equal(t, err, nil)
}

func TestGetHistoricalRateByDateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockIHistoricalRate(ctrl)
	m.EXPECT().GetByDate("").Return(nil, errors.New("test error"))
	_, err := RateUseCase{m}.GetHistoricalRateByDate("")
	assert.NotNil(t, err)
}

func TestGetAnalyze(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockIHistoricalRate(ctrl)
	var analyzes []domain.AnalyzeReport
	analyzes = append(analyzes, domain.AnalyzeReport{
		Currency: "BDT",
		Min:      1,
		Max:      3,
		Avg:      2,
	})
	result := make(map[string]domain.AnalyzeReportResponse)
	for _, value := range analyzes {
		result[value.Currency] = domain.AnalyzeReportResponse{
			Min: value.Min,
			Max: value.Max,
			Avg: value.Avg,
		}
	}
	m.EXPECT().GetAnalyze().Return(analyzes, nil)
	res, err := RateUseCase{m}.GetHistoricalAnalyzes()
	assert.Equal(t, res, result)
	assert.Equal(t, err, nil)
}

func TestGetAnalyzeError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockIHistoricalRate(ctrl)
	m.EXPECT().GetAnalyze().Return(nil, errors.New("test error"))
	_, err := RateUseCase{m}.GetHistoricalAnalyzes()
	assert.NotNil(t, err)
}
