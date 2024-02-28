package entities

type CustomFieldValue struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type CustomFieldValues []CustomFieldValue

func (cfv *CustomFieldValue) TableName() string {
	return "CustomFieldValues"
}
