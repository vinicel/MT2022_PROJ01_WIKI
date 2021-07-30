package connector

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var Db *gorm.DB = initDb()

func initDb() *gorm.DB  {
	fmt.Println(os.Getenv("DATABASE_URL"))
	dsn := fmt.Sprintf("%v:%v@tcp(%v:3306)/wikiGo?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD") ,os.Getenv("DATABASE_URL"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
