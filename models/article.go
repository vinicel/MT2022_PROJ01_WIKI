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
	ID 		  int	 `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type GetOneResponse struct {
	ID        int					`json:"id"`
	CreatedAt time.Time 			`json:"createdAt"`
	UpdatedAt time.Time 			`json:"updatedAt"`
	Title	  string 				`json:"title"`
	Content	  string				`json:"content"`
	AuthorID  int					`json:"authorId"`
	Firstname string				`json:"authorName"`
//	Author	  RelationAuthorResponse `json:"author"`
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
	res := GetOneResponse{}
	am.Db = am.Db.Select("a.*, ac.firstname")
	am.Db = am.Db.Where("a.id", 8)
	am.Db = am.Db.Joins("INNER JOIN accounts ac on ac.id = a.author_id")
	am.Db.Table("articles a").Find(&res)
	return res, nil
}

