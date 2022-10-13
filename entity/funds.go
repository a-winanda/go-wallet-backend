package entity

type Fund struct {
	ID         int    `gorm:"primaryKey;column:id"`
	SourceName string `gorm:"column:source_name"`
}
