package transaction

import (
	"context"

	"github.com/kendoow/gomastering/models"
	"gorm.io/gorm"
)

// TransferTxParams contains the input params of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    models.Transfer `json:"transfer"`
	FromAccount models.Account  `json:"from_account"`
	ToAccount   models.Account  `json:"to_account"`
	FromEntry   models.Entry    `json:"from_entry"`
	ToEntry     models.Entry    `json:"to_entry"`
}

// type of Store
type Store struct {
	db *gorm.DB
}

// Creates new Store
func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

// Creates new transaction
func (s *Store) CreateTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	tx := s.db.Begin()
  
	if tx.Error != nil {
	  return tx.Error
	}
  
	defer func() {
	  if r := recover(); r != nil {
		tx.Rollback()
	  }
	}()
  
	err := fn(tx)
  
	if err != nil {
	  tx.Rollback()
	  return err
	}
  
	return tx.Commit().Error
  }
