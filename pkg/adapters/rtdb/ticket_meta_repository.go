package rtdb

import (
	"fmt"

	"github.com/banyar/lib/pkg/core/common"
	"github.com/banyar/lib/pkg/core/domain/entities"
	"gorm.io/gorm"
)

type TicketMetaRepository struct {
	db *gorm.DB
}

func NewTicketMetaRepository(db *gorm.DB) *TicketMetaRepository {
	return &TicketMetaRepository{
		db: db,
	}
}

func (r *TicketMetaRepository) GetByID(id uint) (*entities.TicketMeta, error) {
	var ticketMeta *entities.TicketMeta

	query := r.db.Select("id", "Status", "Queue", "Created", "LastUpdated").
		Where(fmt.Sprintf(`id = %d `, id)).
		Find(&ticketMeta)

	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, common.ErrRecordNotFound
	}
	return ticketMeta, query.Error
}

func (r *TicketMetaRepository) GetByDatetimeRange(from string, to string, lookupField string) ([]*entities.TicketMeta, error) {
	var ticketMetaCollection []*entities.TicketMeta

	query := r.db.Select("id", "Status", "Queue", "Created", "LastUpdated").
		Where(fmt.Sprintf("%s BETWEEN '%s' AND '%s'", lookupField, from, to)).
		Find(&ticketMetaCollection)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, common.ErrRecordNotFound
	}
	return ticketMetaCollection, query.Error
}
