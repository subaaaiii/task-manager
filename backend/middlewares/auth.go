package middlewares

import (
	"backend/config"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SimpleAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		apiKey := string(config.GetJWTKey())

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization format",
			})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		fmt.Println("TOKEN:", token)
		fmt.Println("APIKEY:", apiKey)

		if token != apiKey {

			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
