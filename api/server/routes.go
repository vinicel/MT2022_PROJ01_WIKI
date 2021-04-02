package server

import (
	"github.com/Wiki-Go/controllers"
	"github.com/Wiki-Go/middleware"
	"github.com/gorilla/mux"
)

func (s *Server) InitialiseRoutes(controller *controllers.Controller) {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/api/v1/login", controller.LoginHandler).Methods("POST")
	s.Router.HandleFunc("/api/v1/accounts", controller.CreateUserHandler).Methods("POST")
	s.Router.HandleFunc("/api/v1/articles/{articleId}/comments", middleware.JWTMiddleware(controller.CreateCommentHandler)).Methods("POST")
	s.Router.HandleFunc("/api/v1/articles/{articleId}/comments", middleware.JWTMiddleware(controller.GetCommentsHandler)).Methods("GET")
	s.Router.HandleFunc("/api/v1/comments/{commentId}", middleware.JWTMiddleware(controller.GetOneCommentHandler)).Methods("GET")
	s.Router.HandleFunc("/api/v1/articles", middleware.JWTMiddleware(controller.GetAllArticles)).Methods("GET")
	s.Router.HandleFunc("/api/v1/articles/{id:[0-9]+}", middleware.JWTMiddleware(controller.GetOne)).Methods("GET")
	s.Router.HandleFunc("/api/v1/articles", middleware.JWTMiddleware(controller.CreateArticle)).Methods("POST")
	s.Router.HandleFunc("/api/v1/articles/{id:[0-9]+}", middleware.JWTMiddleware(controller.UpdateArticle)).Methods("PUT")
}