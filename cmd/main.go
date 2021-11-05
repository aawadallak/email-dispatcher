package main

import (
	"latest/app"
	"latest/app/http"
	"latest/config"
	"latest/config/email"
	"latest/pkg/logger"
)

func main() {

	logger.InitLogger()
	config.Init()
	email.Setup()

	api := http.NewServer()
	api.Run()

	app.Start()
}
