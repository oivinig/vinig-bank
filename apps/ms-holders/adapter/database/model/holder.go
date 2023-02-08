package model

import (
	"time"
)

type Holder struct {
	HolderID   string `gorm:"primary_key; column:holder_id; type:uuid" json:"holder_id"`
	Document   string `gorm:"primary_key; unique" json:"document"`
	HolderName string `gorm:"column:holder_name" json:"holder_name"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
