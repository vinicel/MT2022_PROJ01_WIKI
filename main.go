package main

import (
	"github.com/Wiki-Go/api/server"
	"github.com/Wiki-Go/models"
)

func main()  {

	api := server.Server{
		DB: models.InitGorm(),
	}
	api.Run()
}
