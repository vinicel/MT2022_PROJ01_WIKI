package server

import (
	"github.com/Wiki-Go/controllers"
	"github.com/gorilla/mux"
)

func (s *Server) InitialiseRoutes() {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/api/user/new", controllers.CreateUserHandler).Methods("POST")
}