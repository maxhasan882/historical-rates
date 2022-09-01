package repository

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"github.com/historical-rate/internal/app/adapter/utils"
	"github.com/historical-rate/internal/app/domain"
	"github.com/historical-rate/internal/app/domain/repository"
	"log"
)

const ExternalDataSourceUrl = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"

type DataLoader struct {
	DB *sql.DB
}

func NewDataLoader(db *sql.DB) repository.IDataLoader {
	return &DataLoader{DB: db}
}

// LoadXML returns historical rate from www.ecb.europa.eu.
func (d DataLoader) LoadXML() domain.HistoricalRates {
	api := utils.Api{
		Url:    ExternalDataSourceUrl,
		Method: "GET",
	}
	response, statusCode, err := api.MakeRequest()
	if statusCode != 200 || err != nil {
		log.Println("enable to load data: ", err)
	}
	var data domain.HistoricalRates
	err = xml.Unmarshal(response, &data)
	if err != nil {
		return data
	}
	return data
}

// LoadData save every row of historical rate
func (d DataLoader) LoadData(rate domain.Rate) error {
	_sql := fmt.Sprintf(`INSERT INTO historical_data (date, currency, rate) VALUES ('%s', '%s', '%f');`, rate.Date.Format(DateTimeLayout), rate.Currency, rate.Rate)
	row, err := d.DB.Query(_sql)
	if err != nil {
		return err
	}
	defer func(row *sql.Rows) {
		err = row.Close()
		if err != nil {
			log.Println(err)
		}
	}(row)
	return nil
}
