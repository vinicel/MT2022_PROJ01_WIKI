package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type ArticleModel struct {
	Db		*gorm.DB
}

func (am *ArticleModel) getAll() ([]byte, error) {
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


func (am *ArticleModel) GetOne(id int) (string, error) {
	return `{"toto": "fsd"}`, nil
}
