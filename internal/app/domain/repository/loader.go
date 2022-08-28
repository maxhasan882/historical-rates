package repository

import "github.com/historical-rate/internal/app/domain"

type IDataLoader interface {
	LoadXML() domain.HistoricalRates
	LoadData(rate domain.Rate) error
}
