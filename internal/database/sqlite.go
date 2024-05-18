package database

import (
	"bookkeeping/account"
	"bookkeeping/internal/database/model"
	"bookkeeping/journal"
	"context"
	"gorm.io/gorm"
)

type (
	SQLite struct {
		db *gorm.DB
	}
)

func NewSQLite(db *gorm.DB) *SQLite {
	_ = db.AutoMigrate(
		&model.Transaction{},
		&model.Account{},
		&model.Entry{},
	)

	return &SQLite{
		db: db,
	}
}

func (l SQLite) CreateAccount(a account.Account) error {
	m := model.NewAccount(a)
	if err := l.db.Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func (l SQLite) CreateTransaction(ctx context.Context, incoming, outgoing journal.Entry) error {
	return l.db.Transaction(func(tx *gorm.DB) error {
		trx := model.NewTransaction(outgoing.Transaction)
		if err := tx.WithContext(ctx).Create(&trx).Error; err != nil {
			return err
		}

		from := model.NewEntry(outgoing, trx)
		if err := tx.WithContext(ctx).Create(&from).Error; err != nil {
			return err
		}

		to := model.NewEntry(incoming, trx)
		if err := tx.WithContext(ctx).Create(&to).Error; err != nil {
			return err
		}

		return nil
	})
}
