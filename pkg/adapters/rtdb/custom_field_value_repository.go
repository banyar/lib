package rtdb

import (
	"github.com/banyar/lib/pkg/core/common"
	"github.com/banyar/lib/pkg/core/domain/entities"
	"gorm.io/gorm"
)

type ObjectCustomFieldValuesRepository struct {
	db *gorm.DB
}

/* Creates a new OCFV repository with the given MySQL connection. */
func NewObjectCustomFieldValuesRepository(db *gorm.DB) *ObjectCustomFieldValuesRepository {
	return &ObjectCustomFieldValuesRepository{
		db: db,
	}
}

func (r *ObjectCustomFieldValuesRepository) GetByTicketID(ticketID uint, cfIDList []uint, isDisabled int8) (*entities.ObjectCustomFieldValueCollection, error) {
	ocfvList := &entities.ObjectCustomFieldValueCollection{}
	query := r.db.Where(
		`
		ObjectType = 'RT::Ticket'
		AND ObjectId = ? 
		AND CustomField IN (?) 
		`,
		ticketID, cfIDList)

	if isDisabled < 2 {
		query = query.Where("Disabled = ?", isDisabled)
	}

	query.Find(&ocfvList)

	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, common.ErrRecordNotFound
	}
	return ocfvList, nil
}
