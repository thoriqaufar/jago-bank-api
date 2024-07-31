package repository

import (
	"github.com/jago-bank-api/entity"
	"gorm.io/gorm"
)

type LoginRepository interface {
	EmailExists(email string) bool
	Register(user *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
}

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) *loginRepository {
	return &loginRepository{
		db: db,
	}
}

func (r *loginRepository) EmailExists(email string) bool {
	var user entity.User

	err := r.db.First(&user, "email = ?", email).Error

	return err == nil
}

func (r *loginRepository) Register(user *entity.User) error {
	err := r.db.Create(&user).Error

	mainWallet := entity.Wallet{
		UserID:  user.ID,
		Name:    "Main Wallet",
		Balance: 0,
	}

	err = r.db.Create(&mainWallet).Error

	return err
}

func (r *loginRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := r.db.First(&user, "email = ?", email).Error

	return &user, err
}
