package main

import (
	"log"
	"hotel-booking/config"
	"hotel-booking/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Starting Runers App")

	log.Println("Initializig configuration")
	config := config.InitConfig("runners")

	log.Println("Initializig database")
	dbHandler := server.InitDatabase(config)

	log.Println("Initializig HTTP sever")
	httpServer := server.InitHttpServer(config, dbHandler)

	httpServer.Start()
}