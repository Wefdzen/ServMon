package database

import (
	"fmt"
	"log"

	"github.com/Wefdzen/ServMon/pkg/config"
	"github.com/Wefdzen/ServMon/pkg/db/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Cfg = config.LaunchConfigFile()

func Connect() (*gorm.DB, error) {
	//connect
	urlToDataBase := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", Cfg.PG_user, Cfg.PG_password, Cfg.PG_host, Cfg.PG_port, Cfg.PG_db_name)
	fmt.Println(urlToDataBase)
	db, err := gorm.Open(postgres.Open(urlToDataBase), &gorm.Config{})
	if err != nil {
		log.Fatal("can't open database")
		return nil, err
	}
	db.AutoMigrate(&model.RecordAboutServerInfo{})
	return db, nil
}
