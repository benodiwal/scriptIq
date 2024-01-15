package middlewares

import (
	"net/http"
	"github.com/t01gyl0p/scriptIq/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := utils.TokenValid(ctx)
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}