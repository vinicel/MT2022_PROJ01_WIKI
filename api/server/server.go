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
	//s.DB = models.InitGorm()
	// defer s.DB.Close()
	err := http.ListenAndServe(":8080", s.Router)
	if err != nil {
	 	log.Fatal("ListenAndServe: ", err)
	}

	return s
}
/*
func (s *Server) addDbInHandler(fn func (http.ResponseWriter, *http.Request, *gorm.DB)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, s.DB)
	}
}*/