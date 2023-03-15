package main

import (
	"forum/internal/app"
	"forum/internal/config"
	"log"
)

func main() {
	cfg, err := config.InitConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	app, err := app.InitApp(cfg)
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
}
