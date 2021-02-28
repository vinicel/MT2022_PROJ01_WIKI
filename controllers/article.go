package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Wiki-Go/models"
	"gorm.io/gorm"
	"net/http"
)

func GetAllArticles(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var article models.Article
	type result struct {
		ID		int		`json:"id"`
		Title 	string	`json:"title"`
	}
	var res []result
 	db.Model(&article).Select("ID", "Title").Find(&res)
	fmt.Printf("%v", res)
	output, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
