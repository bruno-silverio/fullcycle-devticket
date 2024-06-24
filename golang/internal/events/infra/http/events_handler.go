package http

import (
	"encoding/json"
	"net/http"

	"github.com/bruno-silverio/fullcycle-devticket/golang/internal/events/usecase"
)

// EventsHandler handles HTTP requests for events.
type EventsHandler struct {
	listEventsUseCase  *usecase.ListEventsUseCase
	listSpotsUseCase   *usecase.ListSpotsUseCase
	getEventUseCase    *usecase.GetEventUseCase
	buyTicketsUseCase  *usecase.BuyTicketsUseCase
	//createEventUseCase *usecase.CreateEventUseCase
	//createSpotsUseCase *usecase.CreateSpotsUseCase
}

// NewEventsHandler creates a new EventsHandler.
func NewEventsHandler(
	listEventsUseCase *usecase.ListEventsUseCase,
	listSpotsUseCase *usecase.ListSpotsUseCase,
	getEventUseCase *usecase.GetEventUseCase,
	buyTicketsUseCase *usecase.BuyTicketsUseCase,
	//createEventUseCase *usecase.CreateEventUseCase,
	//createSpotsUseCase *usecase.CreateSpotsUseCase,
) *EventsHandler {
	return &EventsHandler{
		listEventsUseCase:  listEventsUseCase,
		listSpotsUseCase:   listSpotsUseCase,
		getEventUseCase:    getEventUseCase,
		buyTicketsUseCase:  buyTicketsUseCase,
		//createEventUseCase: createEventUseCase,
		//createSpotsUseCase: createSpotsUseCase,
	}
}

// ListEvents handles the request to list all events.
// @Summary List all events
// @Description Get all events with their details
// @Tags Events
// @Accept json
// @Produce json
// @Success 200 {object} usecase.ListEventsOutputDTO
// @Failure 500 {object} string
// @Router /events [get]
func (h *EventsHandler) ListEvents(w http.ResponseWriter, r *http.Request) {
	output, err := h.listEventsUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// GetEvent handles the request to get details of a specific event.
// @Summary Get event details
// @Description Get details of an event by ID
// @Tags Events
// @Accept json
// @Produce json
// @Param eventID path string true "Event ID"
// @Success 200 {object} usecase.GetEventOutputDTO
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /events/{eventID} [get]
func (h *EventsHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	eventID := r.PathValue("eventID")
	input := usecase.GetEventInputDTO{ID: eventID}

	output, err := h.getEventUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}