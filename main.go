package main

import (
	"github.com/Wiki-Go/api/server"
	"github.com/Wiki-Go/models"
)

func main()  {
	models.InitGorm()
	api := server.Server{}
	api.InitialiseRoutes()
	api.Run()

}
