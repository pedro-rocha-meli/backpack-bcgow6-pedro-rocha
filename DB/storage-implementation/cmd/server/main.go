package main

import (
	"github.com/pedro-rocha-meli/backpack-bcgow6-pedro-rocha/DB/storage-implementation/cmd/server/routes"
	"github.com/pedro-rocha-meli/backpack-bcgow6-pedro-rocha/DB/storage-implementation/pkg/db"
	"log"
)

func main() {
	engine, dataBase := db.ConnectDatabase()
	router := routes.NewRouter(engine, dataBase)
	router.MapRoutes()
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal("Couldn't run engine")
	}
}
