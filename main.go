package main

import (
	"api/app"
	"api/logger"
)

func main() {
	logger.Info("starting application...")
	app.Start()
}
