package main

import (
	"github.com/vinicel/MT2022_PROJ01_WIKI/connector"
	server2 "github.com/vinicel/MT2022_PROJ01_WIKI/server"
	"log"
)


func main()  {
	var err error
	connector.Db, err = connector.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	api := server2.Server{}
	//api.InitialiseRoutes()
	api.Run()
}
