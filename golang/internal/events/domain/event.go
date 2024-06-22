package domain

import (
	"errors"
	"time"
)

var (
	ErrEventNotFound = errors.New("event not found")
	ErrEventNameRequired = errors.New("event name is required")
	ErrEventDateFuture = errors.New("event date must be in the future")
	ErrEventCapacityZero = errors.New("event capacity must be greater than zero")
	ErrEventPriceZero = errors.New("event price must be greater than zero")
)

// Rating represents the age restriction for an event
type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "L10"
	Rating12    Rating = "L12"
	Rating14    Rating = "L14"
	Rating16    Rating = "L16"
	Rating18    Rating = "L18"
)

type User struct {
	ID    string
	Email string
}

// Event represents an event with tickets and spots.
type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartnerID    int
	Spots        []Spot
	Tickets      []Ticket
}

// Validate checks if the event data is valid.
func (e *Event) Validate() error {
	if e.Name == "" {
		return ErrEventNameRequired
	}
	if e.Date.Before(time.Now()) {
		return ErrEventDateFuture
	}
	if e.Capacity <= 0 {
		return ErrEventCapacityZero
	}
	if e.Price <= 0 {
		return ErrEventPriceZero
	}
	return nil
}

// AddSpot adds a spot to the event.
func (e *Event) AddSpot(name string) (*Spot, error) {
	spot, err := NewSpot(e, name)
	if err != nil {
		return nil, err
	}
	e.Spots = append(e.Spots, *spot)
	return spot, nil
}