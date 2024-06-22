package repository

import (
	"database/sql"

	"github.com/bruno-silverio/fullcycle-devticket/golang/internal/events/domain"
)

type mysqlEventRepository struct {
	db *sql.DB
}

// NewMysqlEventRepository creates a new MySQL event repository.
//func NewMysqlEventRepository(db *sql.DB) (domain.EventRepository, error) {
//	return &mysqlEventRepository{db: db}, nil
//}

// CreateSpot inserts a new spot into the database.
func (r *mysqlEventRepository) CreateSpot(spot *domain.Spot) error {
	query := `
		INSERT INTO spots (id, event_id, name, status, ticket_id)
		VALUES (?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, spot.ID, spot.EventID, spot.Name, spot.Status, spot.TicketID)
	return err
}