package database

import (
	"github.com/csyezheng/a2fa/internal/models"
)

func (db *Database) CreateAccount(account *models.Account) error {
	return db.db.Create(account).Error
}

func (db *Database) RetrieveFirstAccount(accountName string, userName string) (account models.Account) {
	db.db.Model(models.Account{AccountName: accountName, Username: userName}).First(&account)
	return
}

func (db *Database) ListAccounts(accountName string, userName string) (accounts []models.Account, err error) {
	db.db.Where(&models.Account{AccountName: accountName, Username: userName}).Find(&accounts)
	return
}

func (db *Database) RemoveAccount(accountName string, userName string) error {
	// When querying with struct, GORM will only query with non-zero fields
	// that means if userName is "", it won’t be used to build query conditions
	return db.db.Where(&models.Account{AccountName: accountName, Username: userName}).Delete(&models.Account{}).Error
}

func (db *Database) SaveAccount(account models.Account) error {
	return db.db.Save(&account).Error
}
