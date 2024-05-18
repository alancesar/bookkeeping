package usecase

import (
	"bookkeeping/account"
	"bookkeeping/journal"
	"bookkeeping/transaction"
	"context"
)

type (
	TransactionDatabase interface {
		CreateTransaction(ctx context.Context, incoming, outgoing journal.Entry) error
	}

	CreateTransactionRequest struct {
		From        account.Account
		To          account.Account
		Amount      uint
		Transaction transaction.Transaction
	}

	CreateTransaction struct {
		db TransactionDatabase
	}
)

func NewCreateTransaction(db TransactionDatabase) *CreateTransaction {
	return &CreateTransaction{
		db: db,
	}
}

func (cr CreateTransaction) Execute(ctx context.Context, request CreateTransactionRequest) error {
	incoming := journal.Entry{
		Credit:      request.Amount,
		Account:     request.To,
		Transaction: request.Transaction,
	}

	outgoing := journal.Entry{
		Debit:       request.Amount,
		Account:     request.From,
		Transaction: request.Transaction,
	}

	return cr.db.CreateTransaction(ctx, incoming, outgoing)
}
