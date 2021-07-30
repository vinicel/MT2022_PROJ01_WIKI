package main

import (
	server2 "github.com/vinicel/MT2022_PROJ01_WIKI/server"
)


func main()  {
	api := server2.Server{}
	//api.InitialiseRoutes()
	api.Run()
}
