package entity

type Wallet struct {
	Id      uint   `gorm:"column:id;primaryKey"`
	UserID  uint   `gorm:"column:user_id"`
	Name    string `gorm:"column:name"`
	Balance int    `gorm:"column:balance"`
	User    User   `gorm:"foreignKey:user_id;references:id"`
}
