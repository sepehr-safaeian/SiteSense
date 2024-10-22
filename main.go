package main

import (
	"SiteSense/cmd"
	"SiteSense/common/logger"
)

func main() {
	logger := logger.NewCustomLogger()

	cmd.Execute()

	logger.Warn("salam")
}
