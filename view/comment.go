package view

import (
	"time"
)

type Comment struct {
	ID        	uint
	Title 		string
	Content 	string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

type CommentDetails struct {
	ID        uint
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Article   Article
}