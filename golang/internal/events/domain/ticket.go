package domain

type TicketType string

const (
	TicketTypeFull TicketType = "full" // Full price ticket
	TicketTypeHalf TicketType = "half" // Half price ticket
)

type Ticket struct {
	ID 					string
	EventID 		string
	Spot 				*Spot
	TicketType 	TicketType
	Price				float64
}