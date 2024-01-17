package service

import (
	"app/internal/repository"
	"context"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp repository.RepositoryTicket
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp repository.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {
	// get all the tickets
	ctx := context.Background()
	tickets, err := s.rp.Get(ctx)
	if err != nil {
		return
	}

	// get the total number of tickets
	total = len(tickets)
	return
}

func (s *ServiceTicketDefault) GetTicketsAmountByDestinationCountry(country string) (total int, err error) {
	// get all the tickets
	ctx := context.Background()
	tickets, err := s.rp.GetTicketsByDestinationCountry(ctx, country)
	if err != nil {
		return
	}

	// get the total number of tickets
	total = len(tickets)
	return
}

func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(country string) (total int, err error) {
	total, err = s.GetTotalAmountTickets()
	if err != nil {
		return
	}
	var totalDestination int
	totalDestination, err = s.GetTicketsAmountByDestinationCountry(country)
	if err != nil {
		return
	}

	return (totalDestination * 100) / total, err
}
