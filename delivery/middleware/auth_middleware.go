package middleware

import (
	"loan_tracker/config"
	"loan_tracker/domain"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authType string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}

		token := authHeaderParts[1]
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}

		// validate token
		var claims domain.Claims
		var err error
		if authType == "access" {
			claims = &domain.LoginClaims{Type: "access"}
			err = config.ValidateToken(token, claims)
		} else {
			claims = &domain.LoginClaims{Type: "refresh"}
			err = config.ValidateToken(token, claims)
		}

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
