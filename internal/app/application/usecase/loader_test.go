package usecase

import (
	"github.com/historical-rate/internal/app/domain"
	"github.com/historical-rate/internal/app/domain/repository/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadDataWithValidData(t *testing.T) {
	cubes := []domain.Cube{{
		Time: "2016-01-02",
		Cube: []struct {
			Currency string  `xml:"currency,attr"`
			Rate     float64 `xml:"rate,attr"`
		}{
			{Currency: "BDT", Rate: 1.2},
			{Currency: "BDT", Rate: 1.3},
			{Currency: "BDT", Rate: 1.4},
		},
	},
		{
			Time: "2016-01-01",
			Cube: []struct {
				Currency string  `xml:"currency,attr"`
				Rate     float64 `xml:"rate,attr"`
			}{
				{Currency: "BDT", Rate: 1.2},
				{Currency: "BDT", Rate: 1.3},
				{Currency: "BDT", Rate: 1.4},
			},
		},
	}
	historicalRates := domain.HistoricalRates{
		Cube: struct {
			Cube []domain.Cube `xml:"Cube"`
		}{Cube: cubes},
	}
	repo := mocks.DataLoaderMoc{}.Returns(historicalRates)
	err := LoaderUseCase{LoadRepository: repo}.LoadData()
	assert.Equal(t, err, nil)
}

func TestLoadDataWithUnValidData(t *testing.T) {
	cubes := []domain.Cube{{
		Time: "20-01-02",
		Cube: []struct {
			Currency string  `xml:"currency,attr"`
			Rate     float64 `xml:"rate,attr"`
		}{
			{Currency: "BDT", Rate: 1.2},
			{Currency: "BDT", Rate: 1.3},
			{Currency: "BDT", Rate: 1.4},
		},
	},
		{
			Time: "2016-01-01",
			Cube: []struct {
				Currency string  `xml:"currency,attr"`
				Rate     float64 `xml:"rate,attr"`
			}{
				{Currency: "BDT", Rate: 1.2},
				{Currency: "BDT", Rate: 1.3},
				{Currency: "BDT", Rate: 1.4},
			},
		},
	}
	historicalRates := domain.HistoricalRates{
		Cube: struct {
			Cube []domain.Cube `xml:"Cube"`
		}{Cube: cubes},
	}
	repo := mocks.DataLoaderMoc{}.Returns(historicalRates)
	err := LoaderUseCase{LoadRepository: repo}.LoadData()
	assert.NotNil(t, err)
}
