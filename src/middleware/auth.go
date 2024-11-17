package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/snekussaurier/minban-backend/mod"
)

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func AuthRequried() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.JSON(http.StatusUnauthorized,
				mod.ErrorResponse{Error: "Authorization header missing"})
			context.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			context.JSON(http.StatusUnauthorized,
				mod.ErrorResponse{Error: "Invalid authorization header format"})
			context.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized,
				mod.ErrorResponse{Error: "Invalid or expired token"})
			context.Abort()
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			context.Set("user_id", claims["user_id"])
		} else {
			context.JSON(http.StatusUnauthorized, mod.ErrorResponse{Error: "Invalid token claims"})
			context.Abort()
		}

		context.Next()
	}
}
