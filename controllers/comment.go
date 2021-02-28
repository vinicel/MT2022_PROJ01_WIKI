package controllers

import (
	"encoding/json"
	"github.com/Wiki-Go/models"
	"io/ioutil"
	"net/http"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
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

	db := models.InitGorm()
	comment := &models.Comment{Title: commentDto.Title, Content: commentDto.Content}
	db.Create(comment)

	output, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(output)
}