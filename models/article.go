package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type ArticleModel struct {
	Db		*gorm.DB
}

type RelationAuthorResponse struct {
	ID 		  int
	Firstname string
	Lastname  string
}

type GetOneResponse struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Title	  string
	Content	  string
	AuthorID  int
	Author	  RelationAuthorResponse
}

type CreateArticleDto struct {
	Title string
	Content string
}

func (am *ArticleModel) GetAll() ([]byte, error) {
	type result struct {
		ID		int		`json:"id"`
		Title	string	`json:"title"`
	}
	var res []result
	modelArticle := Article{}
	am.Db.Model(modelArticle).Association("Accounts").Find(&res)
	toJson, err := json.Marshal(res)
	if err != err {
		return nil, err
	}
	return toJson, nil
}

func (am *ArticleModel) GetOne(id int) (GetOneResponse, error) {
	article := Article{}
	res := GetOneResponse{}
	am.Db.Model(article).Where("id = ?", id).Find(&res)
	return res, nil
}

