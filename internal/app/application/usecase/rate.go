package usecase

import (
	"github.com/historical-rate/internal/app/domain"
	"github.com/historical-rate/internal/app/domain/repository"
)

type RateUseCase struct {
	RateRepository repository.IHistoricalRate
}

func (r RateUseCase) GetLatestHistoricalRate() (map[string]float64, error) {
	rates, err := r.RateRepository.GetLatest()
	if err != nil {
		return nil, err
	}
	result := make(map[string]float64)
	for _, value := range rates {
		result[value.Currency] = value.Rate
	}
	return result, nil
}

func (r RateUseCase) GetHistoricalRateByDate(date string) (map[string]float64, error) {
	rates, err := r.RateRepository.GetByDate(date)
	if err != nil {
		return nil, err
	}
	result := make(map[string]float64)
	for _, value := range rates {
		result[value.Currency] = value.Rate
	}
	return result, nil
}

func (r RateUseCase) GetHistoricalAnalyzes() (map[string]domain.AnalyzeReportResponse, error) {
	rates, err := r.RateRepository.GetAnalyze()
	if err != nil {
		return nil, err
	}
	result := make(map[string]domain.AnalyzeReportResponse)
	for _, value := range rates {
		result[value.Currency] = domain.AnalyzeReportResponse{
			Min: value.Min,
			Max: value.Max,
			Avg: value.Avg,
		}
	}
	return result, nil
}
