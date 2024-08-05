package middleware

import (
	"fmt"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// JWT validation logic
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		log.Println("Authorization Header:", authHeader)

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		tokenString := authParts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// Replace "your-256-bit-secret" with your actual secret key
			return []byte("your_jwt_secret"), nil
		})

		if err != nil {
			log.Println("Token parsing error:", err.Error())
			c.JSON(401, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		if role, err := token.Claims.(jwt.MapClaims); err && token.Valid {
			c.Set("isadmin", role["isadmin"])
		} else {
			c.JSON(401, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, exists := c.Get("isadmin")
		if !exists || !isAdmin.(bool) {
			c.JSON(403, gin.H{"error": "Forbidden: You don't have admin privileges"})
			c.Abort()
			return
		}
		c.Next()
	}
}
