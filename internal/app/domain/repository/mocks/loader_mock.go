package mocks

import (
	"github.com/historical-rate/internal/app/domain"
)

type DataLoaderMoc struct {
	HistoricalRates domain.HistoricalRates
}

func (d DataLoaderMoc) Returns(historicalRates domain.HistoricalRates) DataLoaderMoc {
	d.HistoricalRates = historicalRates
	return d
}

func (d DataLoaderMoc) LoadXML() domain.HistoricalRates {
	historicalRates := d.HistoricalRates
	return historicalRates
}

func (d DataLoaderMoc) LoadData(rate domain.Rate) error {
	return nil
}
