package service

import (
	"context"

	"myproject/internal/model"
)

type (
	ITicketItem interface {
		GetTicketDetailById(ctx context.Context, id int64) (out model.TicketItemOutput, err error)
	}
	ITicketHome interface {
	}
)

var (
	localTicketItem ITicketItem
	localTicketHome ITicketHome
)

func TicketItem() ITicketItem {
	if localTicketItem == nil {
		panic("implement localTicketItem is not found for interface ITicketItem")
	}
	return localTicketItem
}
func InitTicketItem(i ITicketItem) {
	localTicketItem = i
}
func TicketHome() ITicketHome {
	if localTicketHome == nil {
		panic("implement localTicketHome is not found for interface ITicketHome")
	}
	return localTicketHome
}
func InitTicketHome(i ITicketHome) {
	localTicketHome = i
}
