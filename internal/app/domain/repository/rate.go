package repository

import "github.com/historical-rate/internal/app/domain"

type IHistoricalRate interface {
	GetLatest() ([]domain.Rate, error)
	GetByDate(date string) ([]domain.Rate, error)
	GetAnalyze() ([]domain.AnalyzeReport, error)
}
