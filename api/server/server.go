package server

import (
	"github.com/Wiki-Go/controllers"
	"github.com/Wiki-Go/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Router 	*mux.Router
	Logger 	*log.Logger
}

func (s *Server) Run() *Server {
	file, fileErr := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	s.Logger = log.New(file, "logger: ", log.LstdFlags)
	s.Logger.Print("server start running")

	controller := &controllers.Controller{
		Db: models.InitGorm(),
		Logger: s.Logger,
	}
	s.InitialiseRoutes(controller)
	// defer s.DB.Close()
	err := http.ListenAndServe(":8080", s.Router)
	if err != nil {
	 	log.Fatal("ListenAndServe: ", err)
	}

	return s
}