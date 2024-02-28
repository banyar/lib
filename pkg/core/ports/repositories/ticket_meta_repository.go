package repositories

import "github.com/banyar/lib/pkg/core/domain/entities"

type TicketMetaRepository interface {
	/*
		Returns a collection of TicketMeta filtered by Ticket ID

		Example usage:
			result, err := ocfvRepo.GetByID(1234)
	*/
	GetByID(id uint) (*entities.TicketMeta, error)
	/*
		Returns a collection of TicketMeta filtered by Datetime Range

		Pass 'lookupField' as "Created" if you want to filter by ticket created datetime,
		else pass "LastUpdated" to filter by ticket's last updated datetime.

		Example usage:
			result, err := ocfvRepo.GetByDatetimeRange("2024-01-01 00:00:01", "2024-02-30 23:59:59", "LastUpdated")
	*/
	GetByDatetimeRange(from string, to string, lookupField string) ([]*entities.TicketMeta, error)
}
