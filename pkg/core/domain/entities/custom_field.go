package entities

type CustomField struct {
	ID         uint   `json:"custom_field_id" gorm:"primaryKey"`
	Name       string `json:"custom_field_name" gorm:"column:Name"`
	Type       string `json:"custom_field_type" gorm:"column:Type"`
	RenderType string `json:"custom_field_render_type" gorm:"column:RenderType"`
	BasedOn    string `json:"custom_field_based_on" gorm:"column:BasedOn"`
	Disabled   bool   `json:"custom_field_disabled" gorm:"column:Disabled"`
}

type CustomFieldCollection []*CustomField

func (c *CustomField) TableName() string {
	return "CustomFields"
}

func (c *CustomFieldCollection) ExtractIDList() []uint {
	var idList []uint
	for _, ocfv := range *c {
		idList = append(idList, ocfv.ID)
	}
	return idList
}
