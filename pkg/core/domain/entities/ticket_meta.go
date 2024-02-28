package entities

type TicketMeta struct {
	ID          uint   `gorm:"column:id"`
	Status      string `gorm:"column:Status"`
	Queue       int    `gorm:"column:Queue"`
	LastUpdated string `gorm:"column:LastUpdated"`
	Created     string `gorm:"column:Created"`
}

type TicketMetaCollection []TicketMeta

func (*TicketMeta) TableName() string {
	return "Tickets"
}

func (t *TicketMetaCollection) ExtractIDCollection() []uint {
	var idList []uint
	for _, ticket := range *t {
		idList = append(idList, ticket.ID)
	}
	return idList
}
