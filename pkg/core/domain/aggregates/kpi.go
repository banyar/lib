package aggregates

type KPI struct {
	TicketId        int    `gorm:"column:ticket_id;primaryKey"`
	TicketCreatedAt string `gorm:"column:ticket_created_at"`
	TicketStatus    string `gorm:"column:status"`
	Field           string `gorm:"column:field"`
	OldValue        string `gorm:"column:old_value;primaryKey"`
	NewValue        string `gorm:"column:new_value;primaryKey"`
	FromDate        string `gorm:"column:from_date;primaryKey"`
	ToDate          string `gorm:"column:to_date;primaryKey"`
	Duration        string `gorm:"column:duration"`
	KpiAttributes
}

type KpiAttributes struct {
	KpiAttribute_1  string `gorm:"column:attribute_1"`
	KpiAttribute_2  string `gorm:"column:attribute_2"`
	KpiAttribute_3  string `gorm:"column:attribute_3"`
	KpiAttribute_4  string `gorm:"column:attribute_4"`
	KpiAttribute_5  string `gorm:"column:attribute_5"`
	KpiAttribute_6  string `gorm:"column:attribute_6"`
	KpiAttribute_7  string `gorm:"column:attribute_7"`
	KpiAttribute_8  string `gorm:"column:attribute_8"`
	KpiAttribute_9  string `gorm:"column:attribute_9"`
	KpiAttribute_10 string `gorm:"column:attribute_10"`
	KpiAttribute_11 string `gorm:"column:attribute_11"`
	KpiAttribute_12 string `gorm:"column:attribute_12"`
	KpiAttribute_13 string `gorm:"column:attribute_13"`
	KpiAttribute_14 string `gorm:"column:attribute_14"`
	KpiAttribute_15 string `gorm:"column:attribute_15"`
}

func (e *KPI) TableName(tablename string) string {
	return tablename
}
