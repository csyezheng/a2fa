package database

import (
	"github.com/csyezheng/a2fa/internal/models"
)

func (db *Database) CreateAccount(account *models.Account) error {
	return db.db.Create(account).Error
}

func (db *Database) RetrieveFirstAccount(name string) (account models.Account) {
	db.db.Model(models.Account{Name: name}).First(&account)
	return
}

func (db *Database) ListAccounts(names []string) (accounts []models.Account, err error) {
	db.db.Find(&accounts, names)
	return
}

func (db *Database) RemoveAccounts(names []string) error {
	return db.db.Delete(&models.Account{}, names).Error
}

func (db *Database) SaveAccount(account models.Account) error {
	return db.db.Save(&account).Error
}
