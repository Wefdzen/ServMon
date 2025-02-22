package database

import (
	"log"
	"time"

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

func (r *GormUserRepository) GetRec(ipServer string) []model.RecordAboutServerInfo {
	var tmp []model.RecordAboutServerInfo
	//oneMonthAgo := time.Now().AddDate(0, -1, 0).Unix()
	twelveHoursAgo := time.Now().Add(-12*time.Hour - 4*time.Minute).Unix()

	//	r.db.Where("time BETWEEN ? AND ?", oneMonthAgo, time.Now().Unix()).

	r.db.Where("time > ?", twelveHoursAgo).
		Where("ip_server = ?", ipServer).
		//Where("time > ?", twelveHoursAgo). // Только записи новее 12 часов 4 минут
		Order("time DESC").
		Limit(12).
		Find(&tmp)

	return tmp
}
