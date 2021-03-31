package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strings"
)

type Controller struct {
	Db 		*gorm.DB
}

func (c *Controller) WriteJson(w http.ResponseWriter, value interface{}) {
	output, err := json.Marshal(value)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(output)
}

func (c *Controller) getUser(r *http.Request) jwt.MapClaims {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	value, _ := extractClaims(tokenString)
	return value
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecret := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		return nil, false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} //else
	log.Printf("Invalid JWT Token")
	return nil, false
}