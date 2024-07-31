package config

import (
	"github.com/jago-bank-api/entity"
	"github.com/jago-bank-api/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDatabase() {
	dsn := "root:@tcp(localhost:3306)/jago_bank?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	err = db.AutoMigrate(&entity.User{}, &entity.Wallet{})
	helper.PanicIfError(err)

	DB = db
}
