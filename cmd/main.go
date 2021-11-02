package main

import (
	"latest/app"
	"latest/app/http"
	"latest/config"
)

func main() {

	config.Init()
	config.InitLogger()

	api := http.NewServer()
	go api.Run()

	app.Start()
}
