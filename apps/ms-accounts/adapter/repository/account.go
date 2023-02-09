package repository

import (
	"encoding/json"

	"github.com/oivinig/vinigbank/apps/ms-accounts/adapter/database"
	"github.com/oivinig/vinigbank/apps/ms-accounts/adapter/database/model"
	"github.com/oivinig/vinigbank/apps/ms-accounts/domain"
)

type Account struct{}

func (a Account) OpenAccount(am *model.Account) domain.Account {
	var account domain.Account
	database.Connect()

	b, _ := json.Marshal(&am)
	json.Unmarshal(b, &account)
	return account
}
