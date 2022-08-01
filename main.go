package main

import (
	"api/app"
	"api/logger"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("TEST"))
	logger.Info("starting application...")
	app.Start()
}
