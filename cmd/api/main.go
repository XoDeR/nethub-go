package main

import (
	"fmt"

	"go.uber.org/zap"
)

const version = "1.0.0"

func main() {
	// test basic output
	fmt.Println("NetHub version: " + version)

	// load config from env vars

	// enable logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// test logger
	logger.Info("NetHub", zap.String("status", "running"))
}
