package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	Router *mux.Router
}

func (s *Server) Run() *Server {
	s.InitialiseRoutes()
	err := http.ListenAndServe(":8084", s.Router)
	if err != nil {
	 	log.Fatal("ListenAndServe: ", err)
	}

	return s
}

