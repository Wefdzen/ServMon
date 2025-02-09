package database

import (
	"log"

	"github.com/Wefdzen/ServMon/pkg/db/model"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository() *GormUserRepository {
	db, err := Connect()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) AddNewRec(newRecord *model.RecordAboutServerInfo) {
	r.db.Create(&model.RecordAboutServerInfo{Time: newRecord.Time, NameService: newRecord.NameService, IpServer: newRecord.IpServer, LoadAvg5Min: newRecord.LoadAvg5Min, Ram: newRecord.Ram, Memory: newRecord.Memory})
}
