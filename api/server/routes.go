package server

import (
	"github.com/Wiki-Go/controllers"
	"github.com/gorilla/mux"
)

func (s *Server) InitialiseRoutes(controller *controllers.Controller) {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/api/user/new", controller.CreateUserHandler).Methods("POST")
	s.Router.HandleFunc("/comments", controller.CreateCommentHandler).Methods("POST")
	s.Router.HandleFunc("/comments", controller.GetCommentsHandler).Methods("GET")
	s.Router.HandleFunc("/comments/{id}", controller.GetOneCommentHandler).Methods("GET")
	s.Router.HandleFunc("/article", controller.CreateArticle).Methods("POST")
	s.Router.HandleFunc("/article/{id}", controller.UpdateArticle).Methods("PUT")
	s.Router.HandleFunc("/articles", controller.GetAllArticles).Methods("GET")
	s.Router.HandleFunc("/articles/{id:[0-9]+}", controller.GetOne).Methods("GET")
}