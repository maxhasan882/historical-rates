package usecase

import "github.com/historical-rate/internal/app/domain/repository"

type RateUseCase struct {
	RateRepository repository.IHistoricalRate
}

func (r RateUseCase) GetLatestHistoricalRate() (map[string]float32, error) {
	rates, err := r.RateRepository.GetLatest()
	if err != nil {
		return nil, err
	}
	result := make(map[string]float32)
	for _, value := range rates {
		result[value.Currency] = value.Rate
	}
	return result, nil
}
