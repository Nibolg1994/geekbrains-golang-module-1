package main

import (
	"fmt"
	"lesson-7/configuration"
	"log"
)

func main() {
	config, err := configuration.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)
}
