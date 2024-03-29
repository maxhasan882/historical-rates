package repository

import (
	"database/sql"
	"fmt"
	"github.com/historical-rate/internal/app/domain"
	"github.com/historical-rate/internal/app/domain/repository"
)

const DateTimeLayout = "2006-01-02"

type HistoricalRate struct {
	DB *sql.DB
}

func NewHistoricalRate(db *sql.DB) repository.IHistoricalRate {
	return &HistoricalRate{DB: db}
}

// GetLatest get data for the latest date and returns
func (h HistoricalRate) GetLatest() ([]domain.Rate, error) {
	var rates []domain.Rate
	query := `SELECT id, date, currency, rate 
				FROM historical_data his
				WHERE date = (SELECT max(date) from historical_data WHERE id = his.id);`
	data, err := h.DB.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		var rate domain.Rate
		err = data.Scan(&rate.Id, &rate.Date, &rate.Currency, &rate.Rate)
		if err != nil {
			return nil, err
		}
		rates = append(rates, rate)
	}
	return rates, nil
}

// GetByDate historical rates for a given date
func (h HistoricalRate) GetByDate(date string) ([]domain.Rate, error) {
	var rates []domain.Rate
	query := fmt.Sprintf(`SELECT id, date, currency, rate 
				FROM historical_data
				WHERE date BETWEEN '%s 00:00:00'::timestamp and '%s 23:59:59'::timestamp;
				`, date, date)
	data, err := h.DB.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		var rate domain.Rate
		err = data.Scan(&rate.Id, &rate.Date, &rate.Currency, &rate.Rate)
		if err != nil {
			return nil, err
		}
		rates = append(rates, rate)
	}
	return rates, nil
}

// GetAnalyze calculates and returns analyze data
func (h HistoricalRate) GetAnalyze() ([]domain.AnalyzeReport, error) {
	var analyzes []domain.AnalyzeReport
	query := `SELECT currency, MIN(rate), MAX(rate), AVG(rate) 
				FROM historical_data GROUP BY currency;`
	data, err := h.DB.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		var analyze domain.AnalyzeReport
		err = data.Scan(&analyze.Currency, &analyze.Min, &analyze.Max, &analyze.Avg)
		if err != nil {
			return nil, err
		}
		analyzes = append(analyzes, analyze)
	}
	return analyzes, nil
}
