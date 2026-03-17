package controllers

import (
	"goshare/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()

	dsn := "goshare_user:@tcp(127.0.0.1:3306)/goshare?charset=utf8mb4&parseTime=True&loc=Asia%2FJakarta"

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Ada error lo bang: " + err.Error())
	}

	DB.AutoMigrate(
		&models.User{},
		&models.File{},
	)

	return DB
}
