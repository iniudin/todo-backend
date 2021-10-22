package main

import (
	"todo-backend/app"
	"todo-backend/config"
	"todo-backend/helper"
	"todo-backend/router"
)

func main() {
	config, err := config.NewCOnfig()
	helper.PanicIfError(err)

	server := app.NewServer(config)
	router.NewRouter(server)
	server.Start(config.Server.Port)
}
