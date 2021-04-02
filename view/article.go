package view

import "time"

type Article struct {
	ID        		uint
	Title 			string
	Content 		string
	AuthorFirstname string
	AuthorLastname  string
	AuthorEmail  	string
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}