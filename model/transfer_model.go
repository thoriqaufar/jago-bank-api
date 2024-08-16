package model

type TransferRequest struct {
	WalletId          uint `json:"wallet_id" validate:"required"`
	UserDestinationId uint `json:"user_destination_id" validate:"required"`
	Amount            int  `json:"amount" validate:"required"`
}

type SendToKafka struct {
	ID                string `json:"id"`
	WalletId          uint   `json:"wallet_id"`
	UserDestinationId uint   `json:"user_destination_id"`
	Amount            int    `json:"amount"`
}
