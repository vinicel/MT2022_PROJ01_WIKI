package view

import (
	"github.com/Wiki-Go/models"
)

func presentArticle(article models.Article) Article {
	return Article{
		ID: article.ID,
		Title: article.Title,
		Content: article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
}

func presentComment(comment models.Comment) Comment {
	return Comment{
		ID: comment.ID,
		Title: comment.Title,
		Content: comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}

func presentCommentDetails(comment models.Comment) CommentDetails {
	return CommentDetails{
		ID: comment.ID,
		Title: comment.Title,
		Content: comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		Article: presentArticle(comment.Article),
	}
}

func presentComments(comments []models.Comment) []Comment {
	var views []Comment
	for _, comment := range comments {
		views = append(views, presentComment(comment))
	}
	return views
}