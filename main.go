package main

import (
	"github.com/vinicel/MT2022_PROJ01_WIKI/api/server"
)

func main()  {
	api := server.Server{}
	//api.InitialiseRoutes()
	api.Run()
}
