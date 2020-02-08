package main

import (
	"go-rest-todo/app"
	"go-rest-todo/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
