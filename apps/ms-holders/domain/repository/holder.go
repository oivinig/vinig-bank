package repository

import (
	"github.com/oivinig/vinigbank/apps/msholders/adapter/database/model"
	"github.com/oivinig/vinigbank/apps/msholders/domain"
)

// IHolder is interface of holder repository
type IHolder interface {
	CreateHolder(hm *model.Holder) domain.Holder
	FetchHolders() []domain.Holder
	EditHolder(hm *model.Holder) domain.Holder
}
