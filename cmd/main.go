package main

import (
	"latest/app"
	"latest/app/http"
	"latest/config"
)

func main() {

	config.Init()
	config.InitLogger()
	config.Logger().Info("Starting service")

	app.Start()

	api := http.NewServer()
	go api.Run()
}
