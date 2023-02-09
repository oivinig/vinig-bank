package factory

import (
	"github.com/oivinig/vinigbank/apps/ms-accounts/adapter/database/model"
	"github.com/oivinig/vinigbank/apps/ms-accounts/domain"
)

type Account struct{}

func (af Account) OpenAccount(a *model.Account) domain.Account {
	account := domain.Account{
		ID:       a.ID,
		Branch:   a.Branch,
		HolderID: a.HolderID,
		Status:   "unlocked",
	}
	return account
}
