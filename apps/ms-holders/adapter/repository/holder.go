package repository

import (
	"encoding/json"

	"github.com/oivinig/vinigbank/apps/msholders/adapter/database"
	"github.com/oivinig/vinigbank/apps/msholders/adapter/database/model"
	"github.com/oivinig/vinigbank/apps/msholders/domain"
	"github.com/oivinig/vinigbank/apps/msholders/domain/factory"
)

type Holder struct{}

func (h Holder) CreateHolder(hm *model.Holder) domain.Holder {
	db := database.Connect()
	factoryHolder := factory.Holder{}

	holder := factoryHolder.CreateHolder(hm)
	if result := db.Create(&holder); result.Error != nil {
		panic(result.Error)
	}
	return holder
}

func (h Holder) FetchHolders() []domain.Holder {
	var hm []model.Holder
	var holders []domain.Holder
	db := database.Connect()

	if result := db.Find(&hm); result.Error != nil {
		panic(result.Error)
	}
	//FIXME: should implementate a presenter layer
	b, _ := json.Marshal(&hm)
	json.Unmarshal(b, &holders)

	return holders
}

func (h Holder) EditHolder(hm *model.Holder) domain.Holder {
	var holder domain.Holder
	db := database.Connect()
	if result := db.Model(&hm).Where(model.Holder{HolderID: hm.HolderID}).Updates(model.Holder{Document: hm.Document, HolderName: hm.HolderName}); result.Error != nil {
		panic(result.Error)
	}

	// FIXME: should implementate a presenter layer
	b, _ := json.Marshal(&hm)
	json.Unmarshal(b, &holder)
	return holder
}
