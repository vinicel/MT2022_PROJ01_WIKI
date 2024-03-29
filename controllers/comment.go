package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vinicel/MT2022_PROJ01_WIKI/connector"
	"github.com/vinicel/MT2022_PROJ01_WIKI/models"
	"github.com/vinicel/MT2022_PROJ01_WIKI/view"
	"io/ioutil"
	"net/http"
)

func (c *Controller) CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var article models.Article
	result := connector.Db.Preload("Author").First(&article, params["articleId"])
	if result.Error != nil {
		c.ErrorHandler(w, result.Error)
		return
	}

	jsonBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		c.ErrorHandler(w, err)
		return
	}

	var commentDto models.CreateCommentDto
	err = json.Unmarshal(jsonBody, &commentDto)
	if err != nil {
		c.ErrorHandler(w, err)
		return
	}

	id, err := c.getUserIdFromToken(r)
	if err != nil {
		c.ErrorHandler(w, err)
		return
	}
	comment := &models.Comment{Title: commentDto.Title, Content: commentDto.Content, Article: article, AccountID: id}
	connector.Db.Create(comment)

	w.WriteHeader(201)
	c.WriteJson(w, view.PresentComment(*comment))
}

func (c *Controller) GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var comments []models.Comment
	connector.Db.Where("article_id = ?", params["articleId"]).Find(&comments)

	c.WriteJson(w, view.PresentComments(comments))
}

func (c *Controller) GetOneCommentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var comment models.Comment
	result := connector.Db.Preload("Account").Preload("Article.Author").First(&comment, params["commentId"])
	if result.Error != nil {
		c.ErrorHandler(w, result.Error)
		return
	}

	c.WriteJson(w, view.PresentCommentDetails(comment))
}