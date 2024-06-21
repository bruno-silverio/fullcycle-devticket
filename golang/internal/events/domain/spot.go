package domain

import (
	"errors"
)

var (
	ErrInvalidSpotNumber       = errors.New("invalid spot number")
	ErrSpotNotFound            = errors.New("spot not found")
	ErrSpotAlreadyReserved     = errors.New("spot already reserved")
	ErrSpotNameRequired		   	 = errors.New("spot name is required")
	ErrSpotNameTwoCharacters	 = errors.New("spot name must be at least 2 characters long")
	ErrSpotNameStartWithLetter = errors.New("spot name must start with a letter")
	ErrSpotNameEndWithNumber	 = errors.New("spot name must end with a number")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold			SpotStatus = "sold"
)

type Spot struct {
	ID 				string
	EventID 	string
	Name 			string
	Status 		SpotStatus
	TicketID 	string
}

// Validate checks if the spot data is valid.
func (s *Spot) Validate() error {
	if len(s.Name) == 0 {
		return ErrSpotNameRequired
	}
	if len(s.Name) < 2 {
		return ErrSpotNameTwoCharacters
	}
	// Validate if the spot name is in the correct format
	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameStartWithLetter
	}
	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotNameEndWithNumber
	}
	return nil
}
