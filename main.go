package main

import (
	"log"
	"microservice1/config"
	"microservice1/database"
	"microservice1/router"

	"github.com/gin-gonic/gin"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatal("unable to start database")
		return
	}

	err = config.Init()
	if err != nil {
		log.Fatal("unable to configure environemnt based configs")
		return
	}

	r := gin.Default()
	router.InitRouter(r)

	err = r.Run(config.Configurations.Port)
	if err != nil {
		log.Fatal("unable to start server")
		return
	}

	log.Println("Server successfully started on server: " + config.Configurations.Port)
}