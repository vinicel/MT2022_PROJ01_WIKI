package models

import (
	"encoding/json"
	"fmt"
	"github.com/vinicel/MT2022_PROJ01_WIKI/connector"
	"time"
)

type ArticleModel struct {
}

type RelationAuthorResponse struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type GetOneResponse struct {
	ID        uint					`json:"id"`
	CreatedAt time.Time 			`json:"createdAt"`
	UpdatedAt time.Time 			`json:"updatedAt"`
	Title	  string 				`json:"title"`
	Content	  string				`json:"content"`
	AuthorID  int					`json:"authorId"`
	Author	  RelationAuthorResponse `json:"author"`
}

func (am *ArticleModel) GetAll() ([]byte, error) {
	type result struct {
		ID		int		`json:"id"`
		Title	string	`json:"title"`
	}
	var res []result
	modelArticle := Article{}
	connector.Db.Model(modelArticle).Find(&res)
	toJson, err := json.Marshal(res)
	if err != err {
		return nil, err
	}
	return toJson, nil
}

func (am *ArticleModel) GetOne(id int) (GetOneResponse, error) {
	var article Article
	err := connector.Db.Model(article).Where("id = ?", id).Preload("Author").First(&article)
	if err.Error != nil {
		fmt.Printf("%v", err)
		return GetOneResponse{}, err.Error
	}
	res := GetOneResponse{
		ID: article.ID,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		Title: article.Title,
		Content: article.Content,
		AuthorID: article.AuthorID,
		Author: RelationAuthorResponse{
			Firstname: article.Author.Firstname,
			Lastname: article.Author.Lastname,
		},
	}
	return res, nil
}
