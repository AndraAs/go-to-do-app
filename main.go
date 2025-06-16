package main

import (
	"os"

	"todo-app/internal"
	"todo-app/v1"
)

func main() {
	// You can also allow setting this with a --log-level flag
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}
	logger := internal.InitLogger(logLevel)
	logger.Info("Starting To Do CLI App")

	store := v1.NewStore()
	v1.RunCLI(store)
}
