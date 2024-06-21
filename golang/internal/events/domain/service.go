package domain

import (
	"errors"
	"fmt"
)

var ErrInvalidQuantity = errors.New("quantity must be greater than zero")

// Apenas para conseguir atrelar um metodo
type spotService struct{}

// NewSpotService creates a new SpotService.
func NewSpotService() *spotService {
	return &spotService{}
}

// GenerateSpots generates the specified number of spots for an event.
func (s *spotService) GenerateSpots(event *Event, quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	for i := range quantity {
		spotName := fmt.Sprintf("%c%d", 'A'+i/10, i%10+1) // Generate spot name like A1, A2, ..., B1, B2, ...
		spot, err := NewSpot(event, spotName)
		if err != nil {
			return err
		}
		event.Spots = append(event.Spots, *spot)
	}

	return nil
}