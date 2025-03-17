package ticket

import (
	"github.com/gin-gonic/gin"
	"myproject/internal/model"
	"myproject/internal/service"
	"myproject/pkg/response"
)

var TicketItem = new(cTicketItem)

type cTicketItem struct {
}

// GetTicketDetailById godoc
// @Summary      GetTicketDetail by user
// @Description	 User get ticket
// @Tags         Ticket mannagment
// @Accept       json
// @Produce      json
// @Param		 payload body model.TicketItemInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket/item/:id [get]
func (c *cTicketItem) GetTicketDetailById(ctx *gin.Context) {
	var params model.TicketItemInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	ticketItem, err := service.TicketItem().GetTicketDetailById(ctx, params.TicketId)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, ticketItem)
}
