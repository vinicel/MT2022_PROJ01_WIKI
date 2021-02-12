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
	log.Fatal(http.ListenAndServe(":8080", s.Router))

	return s
}

