package controller

import (
	"github.com/gin-gonic/gin"
)

type PongController struct {
}

func NewPongController() *PongController {
	return &PongController{}
}
func (p *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "guest")
	uid := c.Query("uid")
	c.JSON(200, gin.H{
		"message": "pong" + name,
		"uid":     uid,
	})
}
