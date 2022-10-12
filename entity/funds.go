package entity

type Fund struct {
	ID         int    `gorm:"column:id"`
	SourceName string `gorm:"column:source_name"`
}
