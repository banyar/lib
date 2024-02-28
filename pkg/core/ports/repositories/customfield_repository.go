package repositories

import "github.com/banyar/lib/pkg/core/domain/entities"

type CustomFieldRepository interface {
	/*
		Returns a collection of CustomFields filtered by IDs

		Pass 'isDisabled' as (1+) if you want to get all values, both disabled and enabled ones,
		else pass (0) for disabled and (1) for enabled

		Example usage:
			result, err := cfRepo.GetByIDList(customFieldIDList, 2)
	*/
	GetByIDList(customFieldIDList []uint, isDisabled int8) (*entities.CustomFieldCollection, error)
	/*
		Returns a collection of CustomFields filtered by names

		Pass 'isDisabled' as (1+) if you want to get all values, both disabled and enabled ones,
		else pass (0) for disabled and (1) for enabled

		Example usage:
			result, err := cfRepo.GetByNameList(customFieldNameList, 2)
	*/
	GetByNameList(customFieldNameList []string, isDisabled int8) (*entities.CustomFieldCollection, error)
}
