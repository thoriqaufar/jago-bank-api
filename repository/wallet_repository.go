package repository

import (
	"github.com/jago-bank-api/entity"
	"github.com/jago-bank-api/helper"
	"github.com/jago-bank-api/model"
	"gorm.io/gorm"
)

type WalletRepository interface {
	CreateWallet(wallet *entity.Wallet) error
	UpdateWallet(wallet *entity.Wallet, id uint, userId uint) error
	DeleteWallet(wallet *entity.Wallet, id uint, userId uint) error
	ShowAllMyWallets(userId uint) ([]*model.ShowAllWalletResponse, error)
}

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *walletRepository {
	return &walletRepository{
		db: db,
	}
}

func (r *walletRepository) CreateWallet(wallet *entity.Wallet) error {
	err := r.db.Create(&wallet).Error

	return err
}

func (r *walletRepository) UpdateWallet(wallet *entity.Wallet, id uint, userId uint) error {
	row := r.db.Model(&entity.Wallet{}).Where("id = ? AND user_id = ?", id, userId).Updates(&wallet)
	if row.RowsAffected == 0 {
		return &helper.InternalServerError{Message: "Wallet ID does not exist"}
	}

	return nil
}

func (r *walletRepository) DeleteWallet(wallet *entity.Wallet, id uint, userId uint) error {
	row := r.db.Model(&entity.Wallet{}).Where("id = ? AND user_id = ?", id, userId).Delete(&wallet)
	if row.RowsAffected == 0 {
		return &helper.InternalServerError{Message: "Wallet ID does not exist"}
	}

	return nil
}

func (r *walletRepository) ShowAllMyWallets(userId uint) ([]*model.ShowAllWalletResponse, error) {
	var response []*model.ShowAllWalletResponse

	err := r.db.Model(&entity.Wallet{}).Where("user_id = ?", userId).Find(&response).Error

	return response, err
}
