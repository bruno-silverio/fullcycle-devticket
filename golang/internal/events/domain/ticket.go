package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrTicketPriceZero = errors.New("ticket price must be greater than zero")
	ErrInvalidTicketType = errors.New("invalid ticket type")
)

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

// IsValidTicketType checks if a ticket type is valid.
func IsValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

// CalculatePrice calculates the price based on the ticket type.
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

// NewTicket creates a new ticket with the given parameters.
func NewTicket(event *Event, spot *Spot, ticketType TicketType) (*Ticket, error) {
	if !IsValidTicketType(ticketType) {
		return nil, ErrInvalidTicketType
	}

	ticket := &Ticket{
		ID:         uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
		TicketType: ticketType,
		Price:      event.Price,
	}
	ticket.CalculatePrice()
	if err := ticket.Validate(); err != nil {
		return nil, err
	}
	return ticket, nil
}