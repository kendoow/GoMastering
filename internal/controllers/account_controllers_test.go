package controllers

import (
	"testing"
	"time"

	"github.com/kendoow/gomastering/config"
	"github.com/kendoow/gomastering/internal/utils"
	"github.com/kendoow/gomastering/models"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var db *gorm.DB


func init(){
	config.Connect()
	db = config.GetBD()
	db.AutoMigrate(&models.Account{})
}

func createRandomAccount(t *testing.T) models.Account{
	arg := models.Account {
		Owner: utils.RandomOwner(),
		Balance: utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account := arg.CreateAccount()

	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return *account
}

func TestCreateAccount(t *testing.T){
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T)  {
	account1 := createRandomAccount(t)
	account2, _ :=  models.GetAccountById(account1.ID)
	
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}


func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := models.UpdateAccountParams{
		ID:      account1.ID,
		Balance: utils.RandomMoney(),
	}
	
	account2 := models.UpdateAccount(&arg)

	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, _ :=  models.GetAccountById(account1.ID)
	deletedAcc := models.DeleteAccount(account2.ID)
	require.Empty(t, deletedAcc)
}
