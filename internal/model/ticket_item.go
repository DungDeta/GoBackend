package model

type TicketItemOutput struct {
	TicketId       int64  `json:"ticket_id"`
	TicketName     string `json:"ticket_name"`
	StockAvailable int64  `json:"stock_available"`
	StockInitial   int64  `json:"stock_initial"`
}
type TicketItemInput struct {
	TicketId int64 `json:"ticket_id"`
}
