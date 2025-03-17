package impl

import (
	"context"
	"fmt"

	"myproject/internal/database"
	"myproject/internal/model"
)

type sTicketItem struct {
	r *database.Queries
}

func NewTicketItemImpl(r *database.Queries) *sTicketItem {
	return &sTicketItem{
		r: r,
	}
}

func (s *sTicketItem) GetTicketDetailById(ctx context.Context, id int64) (out model.TicketItemOutput, err error) {
	fmt.Println("Service Get TicketDetailById")
	ticket, err := s.r.GetTicketItemById(ctx, id)
	if err != nil {
		return out, err
	}
	out = model.TicketItemOutput{
		TicketId:       ticket.ID,
		TicketName:     ticket.Name,
		StockAvailable: int64(ticket.StockAvailable),
		StockInitial:   int64(ticket.StockInitial),
	}
	return out, nil
}
