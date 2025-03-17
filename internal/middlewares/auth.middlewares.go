package middlewares

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"myproject/internal/utils/auth"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Request url path
		uri := c.Request.URL.Path
		log.Printf("Request URI: %s", uri)
		// Check header
		jwtToken, exists := auth.ExtractBearerToken(c)
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "Unauthorized", "msg": "Missing Authorization Header"})
			return
		}
		// Check token
		claims, err := auth.VerifyTokenSubject(jwtToken)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "Unauthorized", "msg": "Invalid Token"})
			return
		}
		// Set claims to context
		log.Printf("Claims::UUID %v", claims.Subject)
		ctx := context.WithValue(c.Request.Context(), "subjectUUID", claims.Subject)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
