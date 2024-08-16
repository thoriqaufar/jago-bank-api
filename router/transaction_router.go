package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jago-bank-api/config"
	"github.com/jago-bank-api/controller"
	"github.com/jago-bank-api/middleware"
	"github.com/jago-bank-api/repository"
	"github.com/jago-bank-api/service"
)

func TransactionRouter(api *gin.RouterGroup) {
	validate := validator.New()

	transactionRepository := repository.NewTransactionRepository(config.DB)
	transactionService := service.NewTransactionService(transactionRepository, validate)
	transactionController := controller.NewTransactionController(transactionService)

	r := api.Group("/transaction")

	r.Use(middleware.JWTMiddleware())

	r.POST("/transfer", transactionController.Transfer)
}
