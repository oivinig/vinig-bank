package usecase

import (
	"github.com/oivinig/vinigbank/apps/msholders/domain"
	"github.com/oivinig/vinigbank/apps/msholders/domain/repository"
)

func FetchHolders(h repository.IHolder) []domain.Holder {
	holders := h.FetchHolders()
	return holders
}
