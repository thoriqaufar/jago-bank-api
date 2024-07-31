package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/jago-bank-api/entity"
	"github.com/jago-bank-api/helper"
	"github.com/jago-bank-api/model"
	"github.com/jago-bank-api/repository"
)

type WalletService interface {
	CreateWallet(request *model.CreateWalletRequest) error
	UpdateWallet(request *model.UpdateWalletRequest, id uint, userId uint) error
	DeleteWallet(id uint, userId uint) error
	ShowAllMyWallets(userId uint) ([]*model.ShowAllWalletResponse, error)
}

type walletService struct {
	repository repository.WalletRepository
	validate   *validator.Validate
}

func NewWalletService(r repository.WalletRepository, v *validator.Validate) *walletService {
	return &walletService{
		repository: r,
		validate:   v,
	}
}

func (s *walletService) CreateWallet(request *model.CreateWalletRequest) error {
	if err := s.validate.Struct(request); err != nil {
		return &helper.BadRequestError{Message: err.Error()}
	}

	wallet := entity.Wallet{
		UserID:  request.UserID,
		Name:    request.Name,
		Balance: 0,
	}

	if err := s.repository.CreateWallet(&wallet); err != nil {
		return &helper.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *walletService) UpdateWallet(request *model.UpdateWalletRequest, id uint, userId uint) error {
	if err := s.validate.Struct(request); err != nil {
		return &helper.BadRequestError{Message: err.Error()}
	}

	wallet := entity.Wallet{
		Name: request.Name,
	}

	if err := s.repository.UpdateWallet(&wallet, id, userId); err != nil {
		return &helper.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *walletService) DeleteWallet(id uint, userId uint) error {
	var wallet entity.Wallet

	if err := s.repository.DeleteWallet(&wallet, id, userId); err != nil {
		return &helper.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *walletService) ShowAllMyWallets(userId uint) ([]*model.ShowAllWalletResponse, error) {
	allMyWallets, err := s.repository.ShowAllMyWallets(userId)
	if err != nil {
		return nil, &helper.InternalServerError{Message: err.Error()}
	}

	return allMyWallets, nil
}
