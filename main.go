package main

import (
	"SiteSense/cmd"
	"SiteSense/pkg/logger"
)

func main() {
	logger := logger.NewCustomLogger()

	cmd.Execute()

	logger.Info("start using application")
}
