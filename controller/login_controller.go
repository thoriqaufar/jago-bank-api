package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jago-bank-api/helper"
	"github.com/jago-bank-api/model"
	"github.com/jago-bank-api/service"
	"net/http"
)

type loginController struct {
	service service.LoginService
}

func NewLoginController(s service.LoginService) *loginController {
	return &loginController{
		service: s,
	}
}

func (c *loginController) Register(ctx *gin.Context) {
	var register model.RegisterRequest

	if err := ctx.ShouldBindJSON(&register); err != nil {
		helper.HandleError(ctx, &helper.BadRequestError{Message: err.Error()})
		return
	}

	if err := c.service.Register(&register); err != nil {
		helper.HandleError(ctx, err)
		return
	}

	response := model.Response(model.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Account created successfully",
	})

	ctx.JSON(http.StatusCreated, response)
}

func (c *loginController) Login(ctx *gin.Context) {
	var login model.LoginRequest

	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		helper.HandleError(ctx, &helper.BadRequestError{Message: err.Error()})
		return
	}

	result, err := c.service.Login(&login)
	if err != nil {
		helper.HandleError(ctx, err)
		return
	}

	response := model.Response(model.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Login Success",
		Data:       result,
	})

	ctx.JSON(http.StatusOK, response)
}
