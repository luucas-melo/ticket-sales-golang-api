package domain

import (
	"errors"
)

var (
	ErrSpotNameRequired       = errors.New("spot name is required")
	ErrSpotNameTwoChacracters = errors.New(("spot already reserved"))
	ErrSpotStartWithLetter    = errors.New("spot name must start with a letter")
	ErrSpotEndWithNumber      = errors.New("spot name must end with a number")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}

// Validate checks if the spot data is valid.
func (s *Spot) Validate() error {
	if s.Name == "" {
		return ErrSpotNameRequired
	}
	if len(s.Name) < 2 {
		return ErrSpotNameTwoChacracters
	}
	// Validate if the spot name is in the correct format
	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameRequired
	}
	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotEndWithNumber
	}

	return nil
}
