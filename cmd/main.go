package main

import (
	"latest/app/http"
	"latest/config"
	"latest/config/email"
)

func main() {

	config.InitLogger()
	config.Init()
	email.Setup()

	api := http.NewServer()
	api.Run()

	//app.Start()
}
