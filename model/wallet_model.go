package model

type CreateWalletRequest struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name" validate:"required"`
}

type UpdateWalletRequest struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Name   string `json:"name" validate:"required"`
}

type ShowAllWalletResponse struct {
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}
