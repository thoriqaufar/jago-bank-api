package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jago-bank-api/config"
	"github.com/jago-bank-api/router"
)

func main() {
	config.NewDatabase()

	r := gin.Default()
	api := r.Group("/api")

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Jago Bank API",
		})
	})

	router.LoginRouter(api)
	router.WalletRouter(api)
	router.TransactionRouter(api)

	r.Run(":8090")
}
