package task3

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Account struct {
	ID      uint64
	Name    string
	Balance int64
}

type Transaction struct {
	ID            uint64
	FromAccountId uint64
	ToAccountId   uint64
	Amount        int64
}

func TransactionTest(db *gorm.DB) {
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})
	account1 := Account{
		Balance: 1000,
		Name:    "luke",
	}
	account2 := Account{
		Balance: 1000,
		Name:    "peter",
	}
	db.Create([]*Account{&account1, &account2})
	Transfer(account1.ID, account2.ID, 10, db)
	Transfer(account1.ID, account2.ID, 1011, db)
	// db.Where("1=1").Delete(&Account{})
	// db.Where("1=1").Delete(&Transaction{})
}

func Transfer(from uint64, to uint64, amount int64, db *gorm.DB) bool {
	err := db.Transaction(func(tx *gorm.DB) error {
		fromAccount := &Account{ID: from}
		toAccount := &Account{ID: to}
		tx.First(fromAccount)
		tx.First(toAccount)
		if fromAccount.Balance < amount {
			return errors.New("balance not enough")
		}
		fromAccount.Balance -= amount
		toAccount.Balance += amount
		tx.Model(&fromAccount).Select("balance").Updates(fromAccount)
		tx.Model(&toAccount).Select("balance").Updates(toAccount)
		tran := Transaction{
			FromAccountId: from,
			ToAccountId:   to,
			Amount:        amount,
		}
		tx.Create(&tran)
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
