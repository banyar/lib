package primaryadapters

import (
	"github.com/banyar/lib/pkg/adapters/rtdb"
	"github.com/banyar/lib/pkg/core/ports/repositories"
	"gorm.io/gorm"
)

type CustomFieldAdapter struct {
	CustomFieldRepo            repositories.CustomFieldRepository
	ObjectCustomFieldValueRepo repositories.ObjectCustomFieldValuesRepository
}

/*
Returns a new OCFV Facade with the given MySQL GORM instance.

EXAMPLE USAGE:

	OCFV := primaryadapters.NewCustomFieldAdapter(db)

	result := ObjectCustomFieldValueRepo.exampleRepo.exampleMethod()
*/
func NewCustomFieldAdapter(db *gorm.DB) *CustomFieldAdapter {
	customFieldRepo := rtdb.NewCustomFieldRepository(db)
	objectCustomFieldValueRepo := rtdb.NewObjectCustomFieldValuesRepository(db)

	return &CustomFieldAdapter{
		CustomFieldRepo:            customFieldRepo,
		ObjectCustomFieldValueRepo: objectCustomFieldValueRepo,
	}
}
