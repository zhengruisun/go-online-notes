package main

import (
	"go-online-notes/pkg/config"
	"go-online-notes/pkg/logger"
	"log"
)

func main() {
	config, err := config.Load(".")
	if err != nil {
		log.Fatalf("Could not load config, %v", err)
	}

	l := logger.New(config.LogLevel)

}
