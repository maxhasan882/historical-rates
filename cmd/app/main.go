package main

import (
	"github.com/historical-rate/internal/app/adapter/controller"
	"net/http"
)

func main() {
	server := controller.GetServer()
	http.HandleFunc("/rates/latest", server.GetLatestHistoricalRate)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
