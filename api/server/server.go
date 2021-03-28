package server

import (
	"github.com/Wiki-Go/controllers"
	"github.com/Wiki-Go/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	Router 	*mux.Router
}

func (s *Server) Run() *Server {
	controller := &controllers.Controller{
		Db: models.InitGorm(),
	}
	s.InitialiseRoutes(controller)
	err := http.ListenAndServe(":8084", s.Router)
	if err != nil {
	 	log.Fatal("ListenAndServe: ", err)
	}

	return s
}

