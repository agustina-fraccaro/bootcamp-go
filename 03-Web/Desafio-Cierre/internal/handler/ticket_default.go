package handler

import (
	"app/internal/service"
	"net/http"

	"github.com/bootcamp-go/web/response"
)

type TicketDefault struct {
	sv service.ServiceTicket
}

func NewTicketDefault(sv service.ServiceTicket) *TicketDefault {
	return &TicketDefault{
		sv: sv,
	}
}

func (t *TicketDefault) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the total number of tickets
		total, err := t.sv.GetTotalAmountTickets()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message":       "OK",
			"total_tickets": total,
		})

		return
	}
}

func (t *TicketDefault) GetTicketsAmountByDestinationCountry(country string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the total number of tickets
		total, err := t.sv.GetTicketsAmountByDestinationCountry(country)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message":       "OK",
			"total_tickets": total,
		})

		return
	}
}
