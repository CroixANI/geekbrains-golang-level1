package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"lesson9/configuration"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err, "unable to get current folder")
	}

	// simulate env variables by
	// load environment variables from .env file
	err = godotenv.Load(currentDir + "/./config.env")
	if err != nil {
		log.Println("error loading .env file")
	}

	sourceTypeInt := flag.Int("source-type", 0, "Source type [0-2]: 0 - env variables, 1 - json, 2 - yaml")
	flag.Parse()
	if *sourceTypeInt < 0 || *sourceTypeInt > 2 {
		log.Fatalf("incorrect source type: %d", *sourceTypeInt)
	}

	conf, err := configuration.GetConfiguration(configuration.SourceConfigType(*sourceTypeInt))
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(conf)
}
