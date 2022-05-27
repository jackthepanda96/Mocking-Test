package config

import (
	"go-clean/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	dsn := "root:@tcp(localhost:3306)/dbplay?charset=utf8mb4&parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
}
