package repository

import (
	"context"
	"log"

	"github.com/o-ga09/tutorial-ec-backend/app/application/transaction"
)

type TransactionManager struct{}

// RunInTransaction implements transaction.TransactionManager.
func (*TransactionManager) RunInTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	log.Printf("transaction")
	return nil
}

func NewTransactionManager() transaction.TransactionManager {
	return &TransactionManager{}
}
