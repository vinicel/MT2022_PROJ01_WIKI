package server

import (
	"github.com/Wiki-Go/controllers"
	"github.com/gorilla/mux"
)

func (s *Server) InitialiseRoutes(controller *controllers.Controller) {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/", controllers.HelloWorldHandler).Methods("GET")
	s.Router.HandleFunc("/comments", controller.CreateCommentHandler).Methods("POST")
	s.Router.HandleFunc("/comments", controller.GetCommentsHandler).Methods("GET")
	s.Router.HandleFunc("/comments/{id}", controller.GetOneCommentHandler).Methods("GET")
}