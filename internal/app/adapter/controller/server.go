package controller

import (
	"database/sql"
	"github.com/historical-rate/internal/app/adapter/db/connections"
	repo "github.com/historical-rate/internal/app/adapter/repository"
	"github.com/historical-rate/internal/app/domain/repository"
)

type Server struct {
	DB             *sql.DB
	RateRepository repository.IHistoricalRate
}

func GetServer() *Server {
	db := connections.Connect()
	return &Server{DB: db, RateRepository: repo.NewHistoricalRate(db)}
}
