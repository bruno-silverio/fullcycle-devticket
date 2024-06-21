package domain

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
