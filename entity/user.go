package entity

type User struct {
	ID          uint     `gorm:"column:id;primaryKey"`
	Name        string   `gorm:"column:name"`
	Email       string   `gorm:"column:email"`
	PhoneNumber string   `gorm:"column:phone_number"`
	Password    string   `gorm:"column:password"`
	PIN         string   `gorm:"column:pin"`
	Address     string   `gorm:"column:address"`
	Province    string   `gorm:"column:province"`
	City        string   `gorm:"column:city"`
	PostalCode  string   `gorm:"column:postal_code"`
	Wallets     []Wallet `gorm:"foreignKey:user_id;references:id"`
}
