package repositories

import (
	"github.com/bondzai/protocraft/core/entities"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type corePostgresRepository struct {
	db *gorm.DB
}

func NewCorePostgresRepository(db *gorm.DB) CoreRepository {
	return &corePostgresRepository{db: db}
}

func (r *corePostgresRepository) InsertCoreData(in *entities.InsertCoreDto) error {
	data := &entities.Core{
		Amount: in.Amount,
	}

	result := r.db.Create(data)

	if result.Error != nil {
		log.Errorf("InsertCockroachData: %v", result.Error)
		return result.Error
	}

	log.Debugf("InsertCockroachData: %v", result.RowsAffected)
	return nil
}
