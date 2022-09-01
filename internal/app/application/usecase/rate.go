package usecase

import (
	"github.com/historical-rate/internal/app/domain"
	"github.com/historical-rate/internal/app/domain/repository"
)

type RateUseCase struct {
	RateRepository repository.IHistoricalRate
}

// GetLatestHistoricalRate returns all the key value of rate and currency for the latest date
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

// GetHistoricalRateByDate returns all the key value of rate and currency for a given date
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

// GetHistoricalAnalyzes returns min, max and avg value for every currency
func (r RateUseCase) GetHistoricalAnalyzes() (map[string]domain.AnalyzeReport, error) {
	rates, err := r.RateRepository.GetAnalyze()
	if err != nil {
		return nil, err
	}
	result := make(map[string]domain.AnalyzeReport)
	for _, value := range rates {
		result[value.Currency] = domain.AnalyzeReport{
			Min: value.Min,
			Max: value.Max,
			Avg: value.Avg,
		}
	}
	return result, nil
}
