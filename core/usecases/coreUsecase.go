package usecases

import "github.com/bondzai/protocraft/core/models"

type CockroachUsecase interface {
	CockroachDataProcessing(in *models.AddCoreData) error
}
