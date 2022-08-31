package mocks

import (
	"github.com/historical-rate/internal/app/domain"
)

type DataLoaderMoc struct {
	HistoricalRates domain.HistoricalRates
	Error           error
}

func (d DataLoaderMoc) Returns(historicalRates domain.HistoricalRates, err error) DataLoaderMoc {
	d.HistoricalRates = historicalRates
	d.Error = err
	return d
}

func (d DataLoaderMoc) LoadXML() domain.HistoricalRates {
	historicalRates := d.HistoricalRates
	return historicalRates
}

func (d DataLoaderMoc) LoadData(rate domain.Rate) error {
	return d.Error
}
