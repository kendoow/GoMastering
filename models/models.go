package models

import (
	"time"

	"github.com/kendoow/gomastering/config"
	"gorm.io/gorm"
)

var db *gorm.DB

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

type CrateAccountParams struct{
	Owner string `gorm:""json:"owner"`
	Balance  int64 `json:"balance"`
	Currency string `json:"currency"`
}
type UpdateAccountParams struct{
	ID int64 `json:"id"`
	Balance int64 `json:"balance"`
}

type CreateEntryParams struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}


type CreateTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func init(){
	config.Connect()
	db = config.GetBD()
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Entry{})
	db.AutoMigrate(&Transfer{})
}


// account func's
func(b * Account) CreateAccount() *Account{
	db.Create(&b)	
	return b
}

func GetAccountById(id int64) (*Account, *gorm.DB){
	var getAccount Account
	db := db.Where("ID=?", id).Find(&getAccount)
	return &getAccount, db
}

func UpdateAccount(args *UpdateAccountParams) Account {
	var NewAccount Account
	db := db.Where("id = ?", args.ID).First(&args)
	
	db = db.Model(&NewAccount).Update("Balance", args.Balance)
	
	return NewAccount
}

func GetAllAccounts() [] Account{
	var Accounts []Account
	db.Find(&Accounts)
	return Accounts
}


func DeleteAccount(ID int64) Account {
	var Account Account
	db.Where("ID=?",ID).Delete(Account)
	return Account
}

// entries func's

func (b * Entry) CreateEntry() *Entry{
	db.Create(&b)	
	return b
}

func GetEntryById(id int64) (*Entry, *gorm.DB){
	var getEntry Entry
	db := db.Where("ID=?", id).Find(&getEntry)
	return &getEntry, db
}

func GetAllEntries() [] Entry{
	var Entries []Entry
	db.Find(&Entries)
	return Entries
}

func DeleteEntry(ID int64) Entry {
	var Entry Entry
	db.Where("ID=?",ID).Delete(Entry)
	return Entry
}


// transfer func's

func (b * Transfer) CreateTransfer() *Transfer{
	db.Create(&b)	
	return b
}

func GetTransferById(id int64) (*Transfer, *gorm.DB){
	var getTransfer Transfer
	db := db.Where("ID=?", id).Find(&getTransfer)
	return &getTransfer, db
}

func GetAllTransfers() [] Transfer{
	var Transfers []Transfer
	db.Find(&Transfers)
	return Transfers
}

func DeleteTransfer(ID int64) Transfer {
	var Transfer Transfer
	db.Where("ID=?",ID).Delete(Transfer)
	return Transfer
}