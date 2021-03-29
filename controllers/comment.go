package controllers

import (
	"encoding/json"
	"github.com/Wiki-Go/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func (c *Controller) CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var article models.Article
	c.Db.First(&article, params["articleId"])

	jsonBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var commentDto models.CreateCommentDto
	err = json.Unmarshal(jsonBody, &commentDto)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	comment := &models.Comment{Title: commentDto.Title, Content: commentDto.Content, Article: article}
	c.Db.Create(comment)

	c.WriteJson(w, comment)
}

func (c *Controller) GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var comments []models.Comment
	c.Db.Where("article_id = ?", params["articleId"]).Find(&comments)

	c.WriteJson(w, comments)
}

func (c *Controller) GetOneCommentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var comment models.Comment
	c.Db.First(&comment, params["commentId"])

	c.WriteJson(w, comment)
}