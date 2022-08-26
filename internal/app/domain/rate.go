package domain

import (
	"encoding/xml"
	"time"
)

type Cube struct {
	Time string `xml:"time,attr"`
	Cube []struct {
		Currency string  `xml:"currency,attr"`
		Rate     float32 `xml:"rate,attr"`
	} `xml:"Cube"`
}

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

type Rate struct {
	Id       int        `json:"id"`
	Date     *time.Time `json:"date"`
	Currency string     `json:"currency"`
	Rate     float32    `json:"rate"`
}
