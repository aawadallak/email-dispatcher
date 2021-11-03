package main

import (
	"latest/app/http"
	"latest/config"
)

func main() {

	config.Init()
	config.InitLogger()

	api := http.NewServer()
	api.Run()

	//app.Start()
}
