package repositories

import (
	"github.com/banyar/lib/pkg/core/domain/entities"
)

type ObjectCustomFieldValuesRepository interface {
	/*
		Returns a collection of OCFVs filtered by Ticket ID (ObjectId)

		Pass 'isDisabled' as (1+) if you want to get all values, both disabled and enabled ones,
		else pass (0) for disabled and (1) for enabled

		Example usage:
			result, err := ocfvRepo.GetByTicketID(ticketID, cfIDList, 2)
	*/
	GetByTicketID(ticketId uint, idList []uint, isDisabled int8) (*entities.ObjectCustomFieldValueCollection, error)
}
