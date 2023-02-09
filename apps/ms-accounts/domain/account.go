package domain

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
	ID       uint        `json:"account_number"`
	Branch   uint        `json:"branch"`
	HolderID string      `json:"holder_id"`
	Type     AccountType `json:"account_type"`
	Status   Status      `json:"account_status"`
}
