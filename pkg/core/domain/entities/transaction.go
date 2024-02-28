package entities

type Transaction struct {
	Id           uint   `gorm:"column:id"`
	Type         string `gorm:"column:Type"`
	Field        string `gorm:"column:Field"`
	OldValue     string `gorm:"column:OldValue"`
	NewValue     string `gorm:"column:NewValue"`
	OldReference int    `gorm:"column:OldReference"`
	NewReference int    `gorm:"column:NewReference"`
	Created      string `gorm:"column:Created"`
}

type TransactionCollection []*Transaction

func (e *Transaction) TableName() string {
	return "Transactions"
}
