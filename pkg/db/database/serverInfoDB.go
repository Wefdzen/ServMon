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

func (r *GormUserRepository) GetRec(ipServer string, countOfRec int) []model.RecordAboutServerInfo {
	var tmp []model.RecordAboutServerInfo
	//oneMonthAgo := time.Now().AddDate(0, -1, 0).Unix()
	HoursAgo := time.Now().Add(-time.Duration(countOfRec)*time.Hour - 4*time.Minute).Unix()

	//	r.db.Where("time BETWEEN ? AND ?", oneMonthAgo, time.Now().Unix()).

	r.db.Where("time > ?", HoursAgo).
		Where("ip_server = ?", ipServer).
		//Where("time > ?", twelveHoursAgo). // Только записи новее 12 часов 4 минут
		Order("time DESC").
		Limit(countOfRec).
		Find(&tmp)

	return tmp
}
