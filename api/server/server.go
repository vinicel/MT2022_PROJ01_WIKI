package server

import (
	"github.com/Wiki-Go/controllers"
	"github.com/Wiki-Go/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Server struct {
	Router 	*mux.Router
	DB		*gorm.DB
}

func (s *Server) Run() *Server {
	controller := &controllers.Controller{
		Db: s.DB,
	}
	s.InitialiseRoutes(controller)
	s.DB = models.InitGorm()
	// defer s.DB.Close()
	err := http.ListenAndServe(":8084", s.Router)
	if err != nil {
	 	log.Fatal("ListenAndServe: ", err)
	}

	return s
}

