package primaryadapters

import (
	"github.com/banyar/lib/pkg/adapters/rtdb"
	"github.com/banyar/lib/pkg/core/ports/repositories"
	"gorm.io/gorm"
)

type TicketAdapter struct {
	TicketMetaRepo repositories.TicketMetaRepository
}

/*
Returns a new Ticket Facade with the given MySQL GORM instance.

EXAMPLE USAGE:

	TICKET := primaryadapters.NewTicketAdapter(db)

	result := TICKET.exampleRepo.exampleMethod()
*/
func NewMySQLTicketAdapter(db *gorm.DB) *TicketAdapter {
	ticketMetaRepo := rtdb.NewTicketMetaRepository(db)

	return &TicketAdapter{
		TicketMetaRepo: ticketMetaRepo,
	}
}
