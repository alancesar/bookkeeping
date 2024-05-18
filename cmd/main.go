package main

import (
	"bookkeeping/account"
	"bookkeeping/internal/database"
	"bookkeeping/transaction"
	"bookkeeping/usecase"
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqliteDB := database.NewSQLite(db)
	uc := usecase.NewCreateTransaction(sqliteDB)

	from := account.Account{
		Name: "Cash in Bank",
	}

	to := account.Account{
		Name: "Fixed Assets",
	}

	trx := transaction.New("New laptop", time.Now())
	if err := uc.Execute(ctx, usecase.CreateTransactionRequest{
		From:        from,
		To:          to,
		Amount:      1500,
		Transaction: trx,
	}); err != nil {
		log.Fatal(err)
	}
}
