package usecase

import (
	"github.com/oivinig/vinigbank/apps/msholders/adapter/database/model"
	"github.com/oivinig/vinigbank/apps/msholders/domain"
	"github.com/oivinig/vinigbank/apps/msholders/domain/repository"
)

func CreateHolder(h repository.IHolder, hm *model.Holder) domain.Holder {
	holder := h.CreateHolder(hm)
	return holder
}
