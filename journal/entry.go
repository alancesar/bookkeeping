package journal

import (
	"bookkeeping/account"
	"bookkeeping/transaction"
	"time"
)

type (
	Entry struct {
		ID          uint
		Credit      uint
		Debit       uint
		Account     account.Account
		Transaction transaction.Transaction
		CreatedAt   time.Time
	}
)
