package main

import (
	"home-24/app"
)

func main() {
	instance := app.Init()
	instance.Serve(":3000")
}
