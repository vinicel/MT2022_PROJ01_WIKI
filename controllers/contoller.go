package controllers

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

type Controller struct {
	Db		*gorm.DB
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