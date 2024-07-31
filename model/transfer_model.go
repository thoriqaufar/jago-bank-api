package model

type TransferRequest struct {
	WalletId          uint `json:"wallet_id" validate:"required"`
	UserDestinationId uint `json:"user_destination_id" validate:"required"`
	Amount            int  `json:"amount" validate:"required"`
}
