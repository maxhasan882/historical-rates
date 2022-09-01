package domain

import (
	"encoding/xml"
	"time"
)

// Cube is a schema of exchange rate for a particular date
type Cube struct {
	Time string `xml:"time,attr"`
	Cube []struct {
		Currency string  `xml:"currency,attr"`
		Rate     float64 `xml:"rate,attr"`
	} `xml:"Cube"`
}

// HistoricalRates schema of full historical rate response
type HistoricalRates struct {
	XMLName xml.Name `xml:"Envelope"`
	Subject string   `xml:"subject"`
	Sender  struct {
		Name string `xml:"name"`
	} `xml:"Sender"`
	Cube struct {
		Cube []Cube `xml:"Cube"`
	} `xml:"Cube"`
}

// Rate is a db schema for representing every row
type Rate struct {
	Id       int        `json:"id"`
	Date     *time.Time `json:"date"`
	Currency string     `json:"currency"`
	Rate     float64    `json:"rate"`
}

// AnalyzeReport is used for represent a currency report
type AnalyzeReport struct {
	Currency string  `json:"currency"`
	Min      float32 `json:"min"`
	Max      float32 `json:"max"`
	Avg      float32 `json:"avg"`
}

// AnalyzeReportResponse is used for represent the response object of currency report
type AnalyzeReportResponse struct {
	Currency string  `json:"currency,omitempty"`
	Min      float32 `json:"min"`
	Max      float32 `json:"max"`
	Avg      float32 `json:"avg"`
}
