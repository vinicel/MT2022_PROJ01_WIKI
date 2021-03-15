package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"log"
)

type Account struct {
	gorm.Model
	Email 		string 	`gorm:"not null;type:varchar(255);"`
	Password 	string	`gorm:"not null;type:varchar(255);"`
	Firstname 	string	`gorm:"not null;type:varchar(30);"`
	Lastname 	string	`gorm:"not null;type:varchar(30);omitempty"`

}

type Article struct {
	gorm.Model
	Title 		string	`gorm:"not null;type:varchar(40);"`
	Content 	string	`gorm:"not null"`
	AuthorID	int
	Author 		Account
}

type Comment struct {
	gorm.Model
	Title 		string	`gorm:"type:varchar(40);"`
	Content 	string	`gorm:"not null;type:varchar(150);"`
	ArticleID	int
	Article 	Article
	AccountID	int
	Account 	Account
}

type ModelInterface interface {
	Create()
	GetOne(id int)
	GetAll()
	Update(id int)
}

func InitGorm() *gorm.DB {
	dsn := "root:admin@tcp(127.0.0.1:3306)/wikiGo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&Account{}, &Article{}, &Comment{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}