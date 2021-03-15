package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Wiki-Go/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (c *Controller) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	var article models.Article
	type result struct {
		ID		int		`json:"id"`
		Title 	string	`json:"title"`
	}
	var res []result
 	c.Db.Model(&article).Select("ID", "Title").Find(&res)
	fmt.Printf("%v", res)
	output, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}


func (c *Controller) GetOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		code := http.StatusBadRequest
		http.Error(w, http.StatusText(code), code)
		log.Println(err)
	}
	articleModel := &models.ArticleModel{
		Db: c.Db,
	}
	article, _ := articleModel.GetOne(id)
	res, err := json.Marshal(article)
	if err != nil {
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}