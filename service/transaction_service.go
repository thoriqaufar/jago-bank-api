package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/jago-bank-api/helper"
	"github.com/jago-bank-api/model"
	"github.com/jago-bank-api/repository"
)

type TransactionService interface {
	Transfer(fromUser uint, request *model.TransferRequest) error
}

type transactionService struct {
	repository repository.TransactionRepository
	validate   *validator.Validate
}

func NewTransactionService(r repository.TransactionRepository, v *validator.Validate) *transactionService {
	return &transactionService{
		repository: r,
		validate:   v,
	}
}

func (s *transactionService) Transfer(fromUser uint, request *model.TransferRequest) error {
	if err := s.validate.Struct(request); err != nil {
		return &helper.BadRequestError{Message: err.Error()}
	}

	balanceCheck, err := s.repository.BalanceCheck(fromUser, request.WalletId)
	if err != nil {
		return &helper.InternalServerError{Message: err.Error()}
	}

	if balanceCheck.Balance < request.Amount {
		return &helper.InternalServerError{Message: "Not enough balance"}
	}

	if err := s.repository.Reduce(fromUser, request.WalletId, request.Amount); err != nil {
		return &helper.InternalServerError{Message: err.Error()}
	}

	if err := s.repository.Adding(request.UserDestinationId, request.Amount); err != nil {
		return &helper.InternalServerError{Message: err.Error()}
	}

	return nil
}
