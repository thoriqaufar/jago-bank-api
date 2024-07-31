package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jago-bank-api/config"
	"github.com/jago-bank-api/controller"
	"github.com/jago-bank-api/repository"
	"github.com/jago-bank-api/service"
)

func LoginRouter(api *gin.RouterGroup) {
	validate := validator.New()

	loginRepository := repository.NewLoginRepository(config.DB)
	loginService := service.NewLoginService(loginRepository, validate)
	loginController := controller.NewLoginController(loginService)

	api.POST("/register", loginController.Register)
	api.POST("/login", loginController.Login)
}
