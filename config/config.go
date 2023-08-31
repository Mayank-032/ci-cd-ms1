package config

import (
	"log"

	"github.com/spf13/viper"
)

var Configurations Config

type Config struct {
	Environment string `json:"environment"`
	Port        string `json:"port"`
	BaseURL     string `json:"base_url"`
	Database    struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		Schema   string `json:"schema"`
	} `json:"database"`
}

func Init() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error: " + err.Error())
		return err
	}
	err = viper.Unmarshal(&Configurations)

	if err != nil {
		log.Println("Error: " + err.Error())
		return err
	}
	return nil
}