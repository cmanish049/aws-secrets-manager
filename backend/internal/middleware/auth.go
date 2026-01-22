package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, ok := c.Request.BasicAuth()
		if !ok {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
			return
		}

		expectedUser := os.Getenv("BASIC_AUTH_USER")
		expectedPass := os.Getenv("BASIC_AUTH_PASS")

		if expectedUser == "" {
			expectedUser = "admin"
		}
		if expectedPass == "" {
			expectedPass = "admin"
		}

		if user != expectedUser || pass != expectedPass {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		c.Next()
	}
}
