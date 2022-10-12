package entity

type Wallet struct {
	WalletNumber int  `gorm:"column:wallet_number"`
	UserID       int  `gorm:"column:user_id"`
	Balance      int  `gorm:"column:balance"`
	User         User `gorm:"foreignKey:UserID; references:ID"`
}
