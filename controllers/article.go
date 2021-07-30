package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vinicel/MT2022_PROJ01_WIKI/connector"
	"github.com/vinicel/MT2022_PROJ01_WIKI/models"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	article, err := articleModel.GetOne(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), 404)
			return
		}
	}
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

func (c *Controller) CreateArticle(w http.ResponseWriter, r *http.Request){
	jsonBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var articleDto models.Article
	err = json.Unmarshal(jsonBody, &articleDto)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	authorId, err := c.getUserIdFromToken(r)
	if err != nil {
		http.Error(w, "Bad Token", 403)
		return
	}
	article := &models.Article{Title: articleDto.Title, Content: articleDto.Content, AuthorID: authorId}
	connector.Db.Create(article)

	c.WriteJson(w, "ok")
}

func (c *Controller) UpdateArticle(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	jsonBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var articleDto models.Article

	err = json.Unmarshal(jsonBody, &articleDto)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	article := &models.Article{Title: articleDto.Title, Content: articleDto.Content}
	c.Db.First(&article, params["id"]).Updates(map[string]interface{}{"title": articleDto.Title, "content": articleDto.Content})
	c.WriteJson(w, article)
}