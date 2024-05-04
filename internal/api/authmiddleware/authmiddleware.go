package authmiddleware

import (
	"MsLetoChat/internal/support/tokenservice"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema = "Bearer "

		authHeader := ctx.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, BearerSchema) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Missing authorization header scheme",
			})
			return
		}

		token := strings.TrimPrefix(authHeader, BearerSchema)

		claims, err := tokenservice.ParseAccessToken(token)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.Set("userID", claims.Id)
	}
}
