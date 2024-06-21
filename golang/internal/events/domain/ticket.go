package domain

import "errors"

var ErrTicketPriceZero = errors.New("ticket price must be greater than zero")

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

// IsValidTicketKind checks if a ticket kind is valid.
func IsValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

// CalculatePrice calculates the price based on the ticket kind.
func (t *Ticket) CalculatePrice() {
	if t.TicketType == TicketTypeHalf {
		t.Price /= 2
	}
}

// Validate checks if the ticket data is valid.
func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrTicketPriceZero
	}
	return nil
}