package main

import "social/internal/app"

const configDir = "./configs"

func main() {
	app.Run(configDir)
}
