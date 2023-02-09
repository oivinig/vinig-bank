package repository

import (
	"github.com/oivinig/vinigbank/apps/ms-accounts/adapter/database/model"
	"github.com/oivinig/vinigbank/apps/ms-accounts/domain"
)

// IAccount is interface of account repository
type IAccount interface {
	OpenAccount(hm *model.Account) domain.Account
}
