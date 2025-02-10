package database

import "github.com/Wefdzen/ServMon/pkg/db/model"

type UserRepository interface {
	AddNewRec(newRecord *model.RecordAboutServerInfo)
	GetRec(ipServer string) []model.RecordAboutServerInfo
}

func AddNewRecord(repo UserRepository, rec *model.RecordAboutServerInfo) {
	repo.AddNewRec(rec)
}

func GetRecordByIp(repo UserRepository, ipServer string) []model.RecordAboutServerInfo {
	return repo.GetRec(ipServer)
}
