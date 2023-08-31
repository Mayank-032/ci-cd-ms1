package database

import (
	"microservice1/config"
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB
func InitDB() error {
	log.Println("Initialising DB")

	db,  err := sql.Open("mysql", generateDSN())
	if err != nil {
		log.Println("Error: " + err.Error())
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Error: " + err.Error())
		return err
	}

	DB = db
	log.Println("Successfully Initialised DB...")
	return nil
}

func generateDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Configurations.Database.Username,
		config.Configurations.Database.Password,
		config.Configurations.Database.Host,
		config.Configurations.Database.Port,
		config.Configurations.Database.Schema,
	)
}