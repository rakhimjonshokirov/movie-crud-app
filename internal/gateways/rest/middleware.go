package rest

import (
	"net/http"
	"rakhimjonshokirov/movie-crud-app/pkg/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware declares cors policy
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			passHeaders(c)
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func passHeaders(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, UPDATE, GET, PUT, PATCH, DELETE")
	c.Header("Access-Control-Allow-Headers", strings.Join(_whiteListHeader, ","))
	c.Header("Access-Control-Max-Age", "3600")
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ParseJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
