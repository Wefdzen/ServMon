package database

import "github.com/Wefdzen/ServMon/pkg/db/model"

type UserRepository interface {
	AddNewRec(newRecord *model.RecordAboutServerInfo)
}

func AddNewRecord(repo UserRepository, rec *model.RecordAboutServerInfo) {
	repo.AddNewRec(rec)
}
