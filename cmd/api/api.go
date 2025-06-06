package main

import (
	"go.uber.org/zap"
)

type config struct {
	addr        string
	apiURL      string
	frontendURL string
}

type application struct {
	config config
	logger *zap.SugaredLogger
}
