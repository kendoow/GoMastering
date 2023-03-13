package models

import(
	"time"
	"gorm.io/gorm"
	"github.com/kendoow/gomastering/config"
)

type Account struct {
	ID int64 `json:"id"`
	Owner string `json:"owner"`
	Balance int64 `json:"balance"`
	Currency string `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Entry struct {
	ID int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// can be negative or positive
	Amount int64 `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfer struct {
	ID int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID int64 `json:"to_account_id"`
	// must be positive
	Amount int64 `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}





var db *gorm.DB

type CrateAccountParams struct{
	Owner string `gorm:""json:"owner"`
	Balance  int64 `json:"balance"`
	Currency string `json:"currency"`
}

func init(){
	config.Connect()
	db = config.GetBD()
	db.AutoMigrate(&Account{})
}

func(b * Account) CreateAccount() *Account{
	db.Create(&b)	
	return b
}

func GetAllAccounts() [] Account{
	var Accounts []Account
	db.Find(&Accounts)
	return Accounts
}

func GetAccountById(ID int64) (*Account, *gorm.DB){
	var getAccount Account
	db := db.Where("ID=?", ID).Find(&getAccount)
	return &getAccount, db
}

func DeleteAccount(ID int64) Account {
	var Account Account
	db.Where("ID=?",ID).Delete(Account)
	return Account
}