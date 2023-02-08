package factory

import (
	"github.com/google/uuid"
	"github.com/oivinig/vinigbank/apps/msholders/adapter/database/model"
	"github.com/oivinig/vinigbank/apps/msholders/domain"
)

type Holder struct{}

func (hf Holder) CreateHolder(m *model.Holder) domain.Holder {
	holder := domain.Holder{
		HolderID:   uuid.New().String(),
		HolderName: m.HolderName,
		Document:   m.Document,
	}
	return holder
}
