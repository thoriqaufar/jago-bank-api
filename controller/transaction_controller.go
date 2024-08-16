package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jago-bank-api/helper"
	"github.com/jago-bank-api/model"
	"github.com/jago-bank-api/service"
	"net/http"
)

type transactionController struct {
	service service.TransactionService
}

func NewTransactionController(s service.TransactionService) *transactionController {
	return &transactionController{
		service: s,
	}
}

func (c *transactionController) Transfer(ctx *gin.Context) {
	var request model.TransferRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		helper.HandleError(ctx, &helper.BadRequestError{Message: err.Error()})
		return
	}

	userId, _ := ctx.Get("userId")

	if err := c.service.Transfer(uint(userId.(int)), &request); err != nil {
		helper.HandleError(ctx, err)
		return
	}

	response := model.Response(model.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Transfer successfully!",
	})

	ctx.JSON(http.StatusOK, response)
}
