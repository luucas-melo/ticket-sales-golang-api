package domain

type TicketType string

const (
	TicketTypeHalf TicketType = "half" // Half-price ticket
	TickerTypeFull TicketType = "full" // Full-price ticket
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}
