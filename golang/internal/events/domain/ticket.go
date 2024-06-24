package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrTicketPriceZero = errors.New("ticket price must be greater than zero")
	ErrInvalidTicketKind = errors.New("invalid ticket kind")
)

type TicketKind string

const (
	TicketKindFull TicketKind = "full" // Full price ticket
	TicketKindHalf TicketKind = "half" // Half price ticket
)

type Ticket struct {
	ID 					string
	EventID 		string
	Spot 				*Spot
	TicketKind 	TicketKind
	Price				float64
}

// IsValidTicketKind checks if a ticket type is valid.
func IsValidTicketKind(ticketKind TicketKind) bool {
	return ticketKind == TicketKindHalf || ticketKind == TicketKindFull
}

// CalculatePrice calculates the price based on the ticket type.
func (t *Ticket) CalculatePrice() {
	if t.TicketKind == TicketKindHalf {
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
func NewTicket(event *Event, spot *Spot, ticketKind TicketKind) (*Ticket, error) {
	if !IsValidTicketKind(ticketKind) {
		return nil, ErrInvalidTicketKind
	}

	ticket := &Ticket{
		ID:         uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
		TicketKind: ticketKind,
		Price:      event.Price,
	}
	ticket.CalculatePrice()
	if err := ticket.Validate(); err != nil {
		return nil, err
	}
	return ticket, nil
}