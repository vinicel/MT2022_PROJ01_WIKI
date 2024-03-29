package server

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/vinicel/MT2022_PROJ01_WIKI/controllers"
	"github.com/vinicel/MT2022_PROJ01_WIKI/models"
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
	models.InitGorm()
	controller := &controllers.Controller{
		Logger: s.Logger,
	}
	s.InitialiseRoutes(controller)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{os.Getenv("FRONT_URL"), os.Getenv("DOCUMENTATION_URL")},
		AllowedHeaders: []string{"Authorization", "Content-Type", "accept"},
		AllowedMethods: []string{"POST", "GET", "PUT"},
		AllowCredentials: true,
		Debug: false,
	})
	handler := c.Handler(s.Router)
	// defer s.DB.Close()
	err := http.ListenAndServe(os.Getenv("PORT"), handler)
	if err != nil {
	 	log.Fatal("ListenAndServe: ", err)
	}

	return s
}