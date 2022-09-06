package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"lesson8/configuration"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err, "Unable to get current folder")
	}

	// load environment variables from .env file
	err = godotenv.Load(currentDir + "/./config.env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	conf, err := configuration.GetConfiguration()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(conf)
}
