package response

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(200, ResponseData{
		Code: code,
		Msg:  msg[code],
		Data: data,
	})

}

func ErrorResponse(c *gin.Context, code int) {
	c.JSON(200, ResponseData{
		Code: code,
		Msg:  msg[code],
		Data: nil,
	})
}
