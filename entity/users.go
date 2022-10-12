package entity

import "time"

type User struct {
	ID        int       `gorm:"primaryKey;column:id"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	ResetCode string    `gorm:"column:reset_code"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
