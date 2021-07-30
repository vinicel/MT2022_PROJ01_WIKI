package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/vinicel/MT2022_PROJ01_WIKI/connector"
	u "github.com/vinicel/MT2022_PROJ01_WIKI/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"strings"
	"time"
)

type Accounts struct {
	ID          int `json:"id"`
	Email 		string `json:"email"`
	Password 	string `json:"password"`
	Firstname 	string `json:"firstname"`
	Lastname 	string `json:"lastname"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}

type JWTToken struct {
	Token string `json:"token"`
}

type Database struct {
	Account Accounts
}

func (accounts *Database) hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes)
}

func (accounts *Database) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(accounts.Account.Password), []byte(password))
	return err == nil
}

func (accounts *Database) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(accounts.Account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	temp := &Accounts{}

	err :=connector.Db.Table("accounts").Where("email = ?", accounts.Account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. please retry"), false
	}

	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user"), false
	}

	return u.Message(false, "Requiremend passed"), true
}

func (accounts *Database) CreateUser() (map[string]interface{}) {
	if resp, ok := accounts.Validate(); !ok {
		return resp
	}

	accounts.Account.Password = accounts.hashPassword(accounts.Account.Password)
	connector.Db.Create(&accounts.Account)

	response := u.Message(true, "Account created")
	response["account"] = accounts

	return response
}

func (accounts *Database) GenerateJWT(id int, email string) (JWTToken, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 1 * 1).Unix(),
		"user_id": int(id),
		"email": email,
	})
	tokenString, err := token.SignedString(signingKey)

	return JWTToken{tokenString}, err
}