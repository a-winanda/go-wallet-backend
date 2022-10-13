package entity

import (
	"time"
)

type Transaction struct {
	//gorm.Model
	ID              int       `gorm:"column:id"`
	WalletNumber    int       `gorm:"column:wallet_number"`
	TransactionType string    `gorm:"column:transaction_type"`
	SourceID        int       `gorm:"column:source_id"`
	TargetID        int       `gorm:"column:target_id"`
	FundID          int       `gorm:"column:fund_id"`
	Amount          int       `gorm:"column:amount"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	Description     string    `gorm:"column:description"`
	Fund            Fund      `gorm:"foreignKey:FundID;references:ID"`
}

type TransactionRequest struct {
	DescriptionRequest string
	SortByEntity       string
	SortOrder          string
	Limit              int
}
