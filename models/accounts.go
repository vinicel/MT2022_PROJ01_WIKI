package models

import (
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

func (accounts *Accounts) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(accounts.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	temp := &Accounts{}

	err := InitGorm().Table("accounts").Where("email = ?", accounts.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. please retry"), false
	}

	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user"), false
	}

	return u.Message(false, "Requiremend passed"), true
}

func (accounts *Accounts) Create() (map[string]interface{})  {
	if resp, ok := accounts.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(accounts.Password), bcrypt.DefaultCost)
	accounts.Password = string(hashedPassword)

	InitGorm().Create(accounts)

	response := u.Message(true, "Account created")
	response["account"] = accounts

	return response
}
