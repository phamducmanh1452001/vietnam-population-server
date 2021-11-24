package main

import (
	"vietnam-population-server/app"
)

func main() {
	app := &app.App{}
	app.Init()
	app.Run(":3000")
}
