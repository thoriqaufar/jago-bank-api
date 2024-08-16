package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jago-bank-api/helper"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			helper.HandleError(ctx, &helper.UnauthorizedError{Message: "Unauthorized"})
			ctx.Abort()
			return
		}

		userId, err := helper.ValidateToken(tokenString)
		if err != nil {
			helper.HandleError(ctx, &helper.UnauthorizedError{Message: err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("userId", *userId)
		ctx.Next()
	}
}
