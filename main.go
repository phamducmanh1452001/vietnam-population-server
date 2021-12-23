package main

import (
	"os"
	"vietnam-population-server/app"
)

func main() {
	app := &app.App{}
	app.Init()

	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "3000"
	}
	app.Run(":" + port)
}
