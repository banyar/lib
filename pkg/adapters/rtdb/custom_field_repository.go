package rtdb

import (
	"github.com/banyar/lib/pkg/core/common"
	"github.com/banyar/lib/pkg/core/domain/entities"
	"gorm.io/gorm"
)

type CustomFieldRepository struct {
	db *gorm.DB
}

/* Creates a new OCFV repository with the given MySQL connection. */
func NewCustomFieldRepository(db *gorm.DB) *CustomFieldRepository {
	return &CustomFieldRepository{
		db: db,
	}
}

func (r *CustomFieldRepository) GetByNameList(customFieldNameList []string, isDisabled int8) (*entities.CustomFieldCollection, error) {
	cfCollection := &entities.CustomFieldCollection{}
	query := r.db.Where(
		`
		LookupType = 'RT::Queue-RT::Ticket' 
		AND Name IN (?)
		`,
		customFieldNameList)

	if isDisabled < 2 {
		query = query.Where("Disabled = ?", isDisabled)
	}
	query.Find(&cfCollection)

	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, common.ErrRecordNotFound
	}
	return cfCollection, nil
}

func (r *CustomFieldRepository) GetByIDList(customFieldNameList []uint, isDisabled int8) (*entities.CustomFieldCollection, error) {
	cfCollection := &entities.CustomFieldCollection{}
	query := r.db.Where(
		`
		LookupType = 'RT::Queue-RT::Ticket' 
		AND id IN (?)
		`,
		customFieldNameList)

	if isDisabled < 2 {
		query = query.Where("Disabled = ?", isDisabled)
	}

	query.Find(&cfCollection)

	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, common.ErrCustomfieldNotFound
	}
	return cfCollection, nil
}

func (r *CustomFieldRepository) GetCustomFieldIdByName(customFieldNames string) (*entities.CustomField, error) {
	customField := &entities.CustomField{}
	query := r.db.Where("Name = ? AND Disabled = 0", customFieldNames).First(&customField)

	if query.Error != nil {
		return nil, query.Error
	}

	if query.RowsAffected == 0 {
		return nil, common.ErrCustomfieldNotFound
	}
	return customField, nil
}

func (r *CustomFieldRepository) Hello() string {
	return "Hello World"
}

func (r *CustomFieldRepository) PyaePhyoAung() string {
	return "Pyae Phyo Aung"
}
