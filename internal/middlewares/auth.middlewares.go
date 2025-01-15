package middlewares

import (
	"github.com/gin-gonic/gin"
	"myproject/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "vaild-token" {
			response.ErrorResponse(c, response.ErrTokenInvalid)
			c.Abort()
			return
		}
		c.Next()
	}
}
