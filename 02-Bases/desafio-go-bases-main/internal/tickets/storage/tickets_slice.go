package storage

import (
	"errors"
	"strconv"
	"strings"

	ticket "github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

const (
	EarlyMorning = "earlyMorning"
	Morning      = "morning"
	Afternoon    = "afternoon"
	Night        = "night"
)

var (
	ErrInvalidTime           = errors.New("Error: Invalid time")
	ErrCantConvertFlightTime = errors.New("Error: could not convert flight time to int")
)

type TicketsSlice struct {
	tickets []ticket.Ticket
}

func NewTicketSlice() *TicketsSlice {
	return &TicketsSlice{
		tickets: []ticket.Ticket{},
	}
}

func (t *TicketsSlice) AddTicket(id int, name string, email string, destination string, country string, flightTime string, price int) {
	(*t).tickets = append((*t).tickets, ticket.Ticket{
		ID:          id,
		Name:        name,
		Email:       email,
		Destination: destination,
		Country:     country,
		FlightTime:  flightTime,
		Price:       price,
	})
}

func (t TicketsSlice) GetTotalTickets(destination string) (int, error) {
	total := 0
	for _, ticket := range t.tickets {
		if strings.Compare(ticket.Destination, destination) == 0 {
			total++
		}
	}
	return total, nil
}

func (t TicketsSlice) AverageDestination(destination string) (int, error) {
	destinationTotal := 0
	for _, ticket := range t.tickets {
		if strings.Compare(ticket.Destination, destination) == 0 {
			destinationTotal++
		}
	}
	return destinationTotal / len(t.tickets), nil
}

func (t TicketsSlice) GetCountByPeriod(time string) (count int, err error) {
	switch time {
	case EarlyMorning:
		count, err = GetTicketsBetweenTimes(t.tickets, 0, 6)
	case Morning:
		count, err = GetTicketsBetweenTimes(t.tickets, 6, 12)
	case Afternoon:
		count, err = GetTicketsBetweenTimes(t.tickets, 13, 18)
	case Night:
		count, err = GetTicketsBetweenTimes(t.tickets, 19, 23)
	default:
		return count, ErrInvalidTime
	}

	if err != nil {
		return count, err
	}

	return count, nil
}

func GetTicketsBetweenTimes(tickets []ticket.Ticket, initialHour int, finalHour int) (int, error) {
	cantTickets := 0

	for _, ticket := range tickets {
		flightTime, err := strconv.Atoi(ticket.FlightTime[0:2])

		if err != nil {
			return 0, ErrCantConvertFlightTime
		}

		if flightTime >= initialHour && flightTime <= finalHour {
			cantTickets++
		}
	}

	return cantTickets, nil
}
