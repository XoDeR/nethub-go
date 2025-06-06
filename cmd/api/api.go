package main

import (
	"go.uber.org/zap"
)

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type config struct {
	addr        string
	apiURL      string
	frontendURL string
	db          dbConfig
}

type application struct {
	config config
	logger *zap.SugaredLogger
}
