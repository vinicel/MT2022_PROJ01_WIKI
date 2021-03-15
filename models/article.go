package models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
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
	Author	  Account
}

func (am *ArticleModel) GetAll() ([]byte, error) {
	type result struct {
		ID		int		`json:"id"`
		Title	string	`json:"title"`
	}
	var res []result
	modelArticle := Article{}
	am.Db.Model(modelArticle).Select("ID", "Title").Find(&res)
	toJson, err := json.Marshal(res)
	if err != err {
		return nil, err
	}
	return toJson, nil
}

func (am *ArticleModel) GetOne(id int) (Article, error) {
	article := Article{}
//	res := GetOneResponse{}
	am.Db.Model(article).Where("id = ?", id).Find(&article)
	return article, nil
}

