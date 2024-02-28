package rest2

import (
	"github.com/banyar/lib/pkg/core/domain/aggregates"
)

type TicketRepositoryREST struct {
	baseURL string
}

func NewTicketRepo(baseURL string) (*TicketRepositoryREST, error) {
	return &TicketRepositoryREST{
		baseURL: baseURL + "/REST/2.0/ticket/",
	}, nil
}

func (tr *TicketRepositoryREST) GetByID(id int) (*aggregates.Ticket, error) {
	return nil, nil
}

func (tr *TicketRepositoryREST) Update(*aggregates.Ticket) error {
	return nil
}
