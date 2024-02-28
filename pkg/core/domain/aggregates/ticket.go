package aggregates

import "github.com/banyar/lib/pkg/core/domain/entities"

type Ticket struct {
	*entities.TicketMeta
	CustomFields map[string]interface{}
}

type TicketCollection []*Ticket

// WIP
