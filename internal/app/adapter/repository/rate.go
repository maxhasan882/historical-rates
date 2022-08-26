package repository

import (
	"database/sql"
	"github.com/historical-rate/internal/app/domain"
	"github.com/historical-rate/internal/app/domain/repository"
)

type HistoricalRate struct {
	DB *sql.DB
}

func NewHistoricalRate(db *sql.DB) repository.IHistoricalRate {
	return &HistoricalRate{DB: db}
}

func (h HistoricalRate) GetLatest() ([]domain.Rate, error) {
	var rates []domain.Rate
	query := `SELECT id, date, currency, rate 
				FROM historical_data his
				WHERE date = (SELECT max(date) from historical_data WHERE id = his.id)`
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
