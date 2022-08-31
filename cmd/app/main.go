package main

import (
	"github.com/historical-rate/cmd/env"
	"github.com/historical-rate/internal/app/adapter"
	"github.com/historical-rate/internal/app/adapter/controller"
	"github.com/historical-rate/internal/app/application/usecase"
	"log"
	"net/http"
	"os"
)

func init() {
	err := env.LoadFile{LoadRepo: env.NewLoader()}.Load(".env")
	if err != nil {
		log.Println(err)
	}
}

func load(server *controller.Server) {
	err := usecase.LoaderUseCase{LoadRepository: server.LoaderRepository}.LoadData()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	router := adapter.New()
	server := controller.GetServer()
	load(server)
	router.GET("/rates/latest", server.GetLatestHistoricalRate)
	router.GET("/rates/:date", server.GetHistoricalRateByDate)
	router.GET("/rates/analyze", server.GetHistoricalAnalyzeReport)
	log.Println("Server started at port ", os.Getenv("SERVER_PORT"))
	err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), router)
	if err != nil {
		panic(err)
	}
}
