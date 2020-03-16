package main

import (
	"Superhero/app"
)

func main() {
	//config := config.GetConfig()

	app := &app.App{}
	app.Initialize()
	//app.Initial(config)

	//a.Run(":8000")
	app.Run(":8899")
}
