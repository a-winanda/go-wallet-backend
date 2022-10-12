package entity

type Transaction struct {
	ID              int    `gorm:"column:id"`
	WalletNumber    int    `gorm:"column:wallet_number"`
	TransactionType string `gorm:"column:transaction_type"`
	SourceID        int    `gorm:"column:source_id"`
	TargetID        int    `gorm:"column:target_id"`
	FundID          string `gorm:"column:fund_id"`
	Amount          int    `gorm:"column:amount"`
	CreatedAt       string `gorm:"column:created_at"`
	Description     string `gorm:"column:description"`
	Fund            Fund   `gorm:"foreignKey:FundID; references:ID"`
}
