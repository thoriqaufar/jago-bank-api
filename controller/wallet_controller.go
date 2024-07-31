package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jago-bank-api/helper"
	"github.com/jago-bank-api/model"
	"github.com/jago-bank-api/service"
	"net/http"
	"strconv"
)

type walletController struct {
	service service.WalletService
}

func NewWalletController(s service.WalletService) *walletController {
	return &walletController{
		service: s,
	}
}

func (c *walletController) CreateWallet(ctx *gin.Context) {
	var newWallet model.CreateWalletRequest

	if err := ctx.ShouldBindJSON(&newWallet); err != nil {
		helper.HandleError(ctx, &helper.BadRequestError{Message: err.Error()})
		return
	}

	userID, _ := ctx.Get("userId")
	newWallet.UserID = uint(userID.(int))

	if err := c.service.CreateWallet(&newWallet); err != nil {
		helper.HandleError(ctx, err)
		return
	}

	response := model.Response(model.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Wallet created succesfully",
	})

	ctx.JSON(http.StatusCreated, response)
}

func (c *walletController) UpdateWallet(ctx *gin.Context) {
	var walletUpdate model.UpdateWalletRequest

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	userID, _ := ctx.Get("userId")

	if err := ctx.ShouldBindJSON(&walletUpdate); err != nil {
		helper.HandleError(ctx, &helper.BadRequestError{Message: err.Error()})
		return
	}

	if err := c.service.UpdateWallet(&walletUpdate, uint(id), uint(userID.(int))); err != nil {
		helper.HandleError(ctx, err)
		return
	}

	response := model.Response(model.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Wallet updated succesfully",
	})

	ctx.JSON(http.StatusOK, response)
}

func (c *walletController) DeleteWallet(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	userID, _ := ctx.Get("userId")

	if err := c.service.DeleteWallet(uint(id), uint(userID.(int))); err != nil {
		helper.HandleError(ctx, err)
		return
	}

	response := model.Response(model.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Wallet deleted succesfully",
	})

	ctx.JSON(http.StatusOK, response)
}

func (c *walletController) ShowAllMyWallets(ctx *gin.Context) {
	userID, _ := ctx.Get("userId")

	allMyWallets, err := c.service.ShowAllMyWallets(uint(userID.(int)))
	if err != nil {
		helper.HandleError(ctx, err)
		return
	}

	response := model.Response(model.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Loaded All My Wallets",
		Data:       allMyWallets,
	})

	ctx.JSON(http.StatusOK, response)
}
