package entities

type ObjectCustomFieldValue struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	TicketID      uint   `json:"ticket_id" gorm:"column:ObjectId"`
	CustomFieldID uint   `json:"customfield" gorm:"column:CustomField"`
	Content       string `json:"content" gorm:"column:Content"`
	Created       string `json:"created" gorm:"column:Created"`
}

type ObjectCustomFieldValueCollection []*ObjectCustomFieldValue

func (o *ObjectCustomFieldValue) TableName() string {
	return "ObjectCustomFieldValues"
}

func (o *ObjectCustomFieldValueCollection) ExtractIDList() []uint {
	var idList []uint
	for _, ocfv := range *o {
		idList = append(idList, ocfv.ID)
	}
	return idList
}
