package model

import (
	"bookkeeping/account"
	"bookkeeping/journal"
	"bookkeeping/transaction"
	"time"
)

type (
	Entry struct {
		ID            uint `gorm:"primarykey"`
		Credit        uint
		Debit         uint
		Account       Account
		AccountID     uint
		Transaction   Transaction
		TransactionID uint
		CreatedAt     time.Time
	}

	Account struct {
		ID        uint `gorm:"primarykey"`
		Name      string
		CreatedAt time.Time
	}

	Transaction struct {
		ID          uint `gorm:"primarykey"`
		Description string
		RequestedAt time.Time
		CreatedAt   time.Time
	}
)

func NewEntry(entry journal.Entry, trx Transaction) Entry {
	return Entry{
		ID:            entry.ID,
		Credit:        entry.Credit,
		Debit:         entry.Debit,
		Account:       NewAccount(entry.Account),
		AccountID:     entry.Account.ID,
		Transaction:   trx,
		TransactionID: trx.ID,
		CreatedAt:     entry.CreatedAt,
	}
}

func NewAccount(a account.Account) Account {
	return Account{
		ID:        a.ID,
		Name:      a.Name,
		CreatedAt: a.CreatedAt,
	}
}

func NewTransaction(t transaction.Transaction) Transaction {
	return Transaction{
		ID:          t.ID,
		Description: t.Description,
		RequestedAt: t.RequestedAt,
		CreatedAt:   t.CreatedAt,
	}
}
