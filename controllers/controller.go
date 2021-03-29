package controllers

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"log"
	"net/http"
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