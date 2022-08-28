package usecase

import (
	"github.com/historical-rate/internal/app/adapter/utils"
	"github.com/historical-rate/internal/app/domain"
	"github.com/historical-rate/internal/app/domain/repository"
)

const DateTimeLayout = "2006-01-02"

type LoaderUseCase struct {
	LoadRepository repository.IDataLoader
}

func (l LoaderUseCase) LoadData() error {
	data := l.LoadRepository.LoadXML()
	for _, data := range data.Cube.Cube {
		dateTime, err := utils.StringToDate(data.Time, DateTimeLayout)
		if err != nil {
			return err
		}
		for _, rt := range data.Cube {
			row := domain.Rate{
				Date:     &dateTime,
				Currency: rt.Currency,
				Rate:     rt.Rate,
			}
			err = l.LoadRepository.LoadData(row)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
