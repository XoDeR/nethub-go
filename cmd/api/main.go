package main

import (
	"fmt"

	"github.com/XoDeR/nethub-go/internal/env"
	"go.uber.org/zap"
)

const version = "1.0.0"

func main() {
	// test basic output
	fmt.Println("NetHub version: " + version)

	// load config from env vars
	cfg := config{
		addr:        env.GetString("ADDR", ":8080"),
		apiURL:      env.GetString("EXTERNAL_URL", "localhost:8080"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:5173"),
	}

	// enable logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// test logger
	logger.Info("NetHub", zap.String("status", "running"))

	app := &application{
		config: cfg,
		logger: logger,
	}

	logger.Info(app)
}
