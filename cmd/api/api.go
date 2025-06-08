package main

import (
	"net/http"

	"github.com/XoDeR/nethub-go/internal/store"
	"github.com/go-chi/chi/v5"
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
	store  store.Storage
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	return r
}

func (app *application) run(mux http.Handler) error {
	return nil
}
