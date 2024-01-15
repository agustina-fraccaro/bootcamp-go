package repository

import (
	"app/internal"
	"context"
)

// RepositoryTicket represents the repository interface for tickets
type RepositoryTicket interface {
	// GetAll returns all the tickets
	Get(ctx context.Context) (t map[int]internal.TicketAttributes, err error)

	// GetTicketByDestinationCountry returns the tickets filtered by destination country
	GetTicketsByDestinationCountry(ctx context.Context, country string) (t map[int]internal.TicketAttributes, err error)
}
