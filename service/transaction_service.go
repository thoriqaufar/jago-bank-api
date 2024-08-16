package service

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

	config := &kafka.ConfigMap{
		//"bootstrap.servers": "localhost:9092",
		"bootstrap.servers": "kafka:9093",
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	topic := "transaction"
	transaction := model.SendToKafka{
		ID:                uuid.New().String(),
		WalletId:          request.WalletId,
		UserDestinationId: request.UserDestinationId,
		Amount:            request.Amount,
	}

	value, err := json.Marshal(transaction)
	if err != nil {
		panic(err)
	}

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: value,
	}

	err = producer.Produce(msg, nil)
	if err != nil {
		panic(err)
	}

	return nil
}
