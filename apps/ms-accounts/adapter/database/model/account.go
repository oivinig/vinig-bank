package model

import "time"

type Status string
type AccountType string
type TransactionType string

const (
	LockedStatus          = "locked"
	UnlockedStatus        = "unlocked"
	CheckingAccountType   = "checking"
	SavingAccountType     = "saving"
	CreditTransactionType = "credit"
	DebitTransactionType  = "debit"
)

type Account struct {
	ID        uint        `gorm:"primary_key; unique; auto_increment; column:account_id"`
	Branch    uint        `gorm:"primary_key;not_null" json:"branch"`
	HolderID  string      `gorm:"primary_key; type:uuid" json:"holder_id"`
	Type      AccountType `gorm:"<-:create, not_null; column:account_type" validate:"nonnil" json:"account_type"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
