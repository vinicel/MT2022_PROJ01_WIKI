package main

import (
	"github.com/vinicel/Wiki-Go/api/server"
)

func main()  {
	api := server.Server{}
	api.InitialiseRoutes()
	api.Run()

}
