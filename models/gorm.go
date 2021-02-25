package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"log"
)

type Account struct {
	gorm.Model
	Email string
	Password string
	Firstname string
	Lastname string

}

type Article struct {
	gorm.Model
	Title string
	Content string

}

type Comment struct {
	gorm.Model
	Title string
	Content string

}

func InitGorm() *gorm.DB {
	dsn := "root:admin@tcp(127.0.0.1:3306)/wikiGo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&Account{})
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(&Article{})
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(&Comment{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}