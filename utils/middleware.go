package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := TokenValid(ctx)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}
		ctx.Next()

	}
}
