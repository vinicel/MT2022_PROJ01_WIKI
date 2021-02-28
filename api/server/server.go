package server

import (
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
	s.InitialiseRoutes()
	s.DB = models.InitGorm()
	// defer s.DB.Close()
	err := http.ListenAndServe(":8084", s.Router)
	if err != nil {
	 	log.Fatal("ListenAndServe: ", err)
	}

	return s
}

func (s *Server) addDbInHandler(fn func (http.ResponseWriter, *http.Request, *gorm.DB)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, s.DB)
	}
}

