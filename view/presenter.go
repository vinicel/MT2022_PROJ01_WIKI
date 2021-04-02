package view

import (
	"wiki/models"
)

func PresentArticle(article models.Article) Article {
	return Article{
		ID: article.ID,
		Title: article.Title,
		Content: article.Content,
		AuthorFirstname: article.Author.Firstname,
		AuthorLastname: article.Author.Lastname,
		AuthorEmail: article.Author.Email,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
}

func PresentComment(comment models.Comment) Comment {
	return Comment{
		ID: comment.ID,
		Title: comment.Title,
		Content: comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}

func PresentCommentDetails(comment models.Comment) CommentDetails {
	return CommentDetails{
		ID: comment.ID,
		Title: comment.Title,
		Content: comment.Content,
		Firstname: comment.Account.Firstname,
		Lastname: comment.Account.Lastname,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		Article: PresentArticle(comment.Article),
	}
}

func PresentComments(comments []models.Comment) []Comment {
	views := []Comment{}
	for _, comment := range comments {
		views = append(views, PresentComment(comment))
	}
	return views
}