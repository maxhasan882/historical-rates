package main

import (
	"github.com/historical-rate/internal/app/adapter"
	"github.com/historical-rate/internal/app/adapter/controller"
	"log"
	"net/http"
	"os"
)

func main() {
	router := adapter.New()
	server := controller.GetServer()
	router.GET("/rates/latest", server.GetLatestHistoricalRate)
	router.GET("/rates/:date", server.GetHistoricalRateByDate)
	router.GET("/rates/analyze", server.GetHistoricalAnalyzeReport)
	log.Println("Server started at port ", os.Getenv("SERVER_PORT"))
	err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), router)
	if err != nil {
		panic(err)
	}
}
