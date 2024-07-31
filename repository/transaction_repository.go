package repository

import (
	"github.com/jago-bank-api/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Adding(toUser uint, amount int) error
	Reduce(fromUser uint, walletId uint, amount int) error
	BalanceCheck(fromUser uint, walletId uint) (*entity.Wallet, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) Adding(toUser uint, amount int) error {
	err := r.db.Model(&entity.Wallet{}).Where("user_id = ? AND name = 'Main Wallet'", toUser).UpdateColumn("balance", gorm.Expr("balance + ?", amount)).Error

	return err
}

func (r *transactionRepository) Reduce(fromUser uint, walletId uint, amount int) error {
	err := r.db.Model(&entity.Wallet{}).Where("user_id = ? AND id = ?", fromUser, walletId).UpdateColumn("balance", gorm.Expr("balance - ?", amount)).Error

	return err
}

func (r *transactionRepository) BalanceCheck(fromUser uint, walletId uint) (*entity.Wallet, error) {
	var wallet entity.Wallet

	err := r.db.Model(&entity.Wallet{}).Where("user_id = ? AND id = ?", fromUser, walletId).First(&wallet).Error

	return &wallet, err
}
