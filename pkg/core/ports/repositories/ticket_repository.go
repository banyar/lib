package repositories

import (
	"github.com/banyar/lib/pkg/core/domain/aggregates"
)

type TicketRepository interface {
	GetByID(id int) (*aggregates.Ticket, error)
	Update(ticket *aggregates.Ticket) error
}
