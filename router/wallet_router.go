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

func WalletRouter(api *gin.RouterGroup) {
	validate := validator.New()

	walletRepository := repository.NewWalletRepository(config.DB)
	walletService := service.NewWalletService(walletRepository, validate)
	walletController := controller.NewWalletController(walletService)

	r := api.Group("/wallet")

	r.Use(middleware.JWTMiddleware())

	r.POST("/create", walletController.CreateWallet)
	r.PUT("/update/:id", walletController.UpdateWallet)
	r.DELETE("/delete/:id", walletController.DeleteWallet)
	r.GET("/", walletController.ShowAllMyWallets)
}
