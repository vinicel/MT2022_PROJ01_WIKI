package models

import (
	"fmt"
	u "github.com/Wiki-Go/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Accounts struct {
	Email 		string `json:"email"`
	Password 	string `json:"password"`
	Firstname 	string `json:"firstname"`
	Lastname 	string `json:"lastname"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}

type Database struct {
	Db 	*gorm.DB
	Account Account
}

func (accounts *Database) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(accounts.Account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	temp := &Accounts{}

	err := InitGorm().Table("accounts").Where("email = ?", accounts.Account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. please retry"), false
	}

	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user"), false
	}

	return u.Message(false, "Requiremend passed"), true
}

func (accounts *Database) CreateUser() (map[string]interface{}) {
/*	if resp, ok := accounts.Validate(); !ok {
		return resp
	}
*/
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(accounts.Account.Password), bcrypt.DefaultCost)

	accounts.Account.Password = string(hashedPassword)
	fmt.Printf("Email: %v\n Password: %v\n Firstname: %v\n Lastname: %v\n", accounts.Account.Email, accounts.Account.Password, accounts.Account.Firstname,accounts.Account.Lastname)
	accounts.Db.Create(&accounts.Account)
	/*if err != nil && err != gorm.ErrRecordNotFound {
		log.Fatal(err)
		return u.Message(false, "Connection error. please retry")
	}*/
	//InitGorm().Create(accounts)

	response := u.Message(true, "Account created")
	response["account"] = accounts

	return response
}
