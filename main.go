package main

import "wiki/api/server"

func main()  {
	api := server.Server{}
	//api.InitialiseRoutes()
	api.Run()
}
