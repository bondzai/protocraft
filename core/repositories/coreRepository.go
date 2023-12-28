package repositories

import "github.com/bondzai/protocraft/core/entities"

type CoreRepository interface {
	InsertCoreData(in *entities.InsertCoreDto) error
}
