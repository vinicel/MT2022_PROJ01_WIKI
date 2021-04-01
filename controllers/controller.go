package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Controller struct {
	Db 		*gorm.DB
	Logger 	*log.Logger
}

func (c *Controller) WriteJson(w http.ResponseWriter, value interface{}) {
	output, err := json.Marshal(value)
	if err != nil {
		c.ErrorHandler(w, err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	_, writeErr := w.Write(output)
	c.ErrorHandler(w, writeErr)
}

func (c *Controller) ErrorHandler(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	c.Logger.Print(err.Error())

	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, err.Error(), 404)
		return
	}

	http.Error(w, http.StatusText(500), 500)
}

func (c *Controller) getUserAuthenticated(r *http.Request) jwt.MapClaims {
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

func (c *Controller) getUserIdFromToken(r *http.Request) (int, error) {
	user := c.getUserAuthenticated(r)
	userId, err := strconv.Atoi(fmt.Sprintf("%v", user["user_id"]))
	if err != nil {
		return -1, err
	}
	return userId, nil
}