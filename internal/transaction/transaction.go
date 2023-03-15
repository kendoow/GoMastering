package transaction

import (
	"context"
	// "time"

	// "github.com/kendoow/gomastering/models"
	"gorm.io/gorm"
)

// TransferTsx performs a money from one account to the other
// It creates a transfer record, add account entries and update accounts' balance within a single database transaction

func(store *Store)TransferTsx(ctx context.Context, arg TransferTxParams)(TransferTxResult, error){
	var result TransferTxResult

	err := store.CreateTransaction(ctx, func(tx *gorm.DB) error {

		// transfer := models.Transfer{
		// 	FromAccountID: arg.FromAccountID,
		// 	ToAccountID: arg.ToAccountID,
		// 	Amount: arg.Amount,

		// }

		


		return nil
	})

	return result, err
}