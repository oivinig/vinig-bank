package usecase

import (
	"github.com/oivinig/vinigbank/apps/ms-accounts/adapter/database/model"
	"github.com/oivinig/vinigbank/apps/ms-accounts/domain"
	"github.com/oivinig/vinigbank/apps/ms-accounts/domain/repository"
)

func OpenAccount(a repository.IAccount, am *model.Account) domain.Account {
	account := a.OpenAccount(am)
	return account
}
