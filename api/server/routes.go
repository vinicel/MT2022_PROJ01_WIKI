package server

import (
	"github.com/gorilla/mux"
	"github.com/vinicel/Wiki-Go/controllers"
)

func (s *Server) InitialiseRoutes() {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/", controllers.HelloWorldHandler).Methods("GET")
}