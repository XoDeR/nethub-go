package main

import (
	"fmt"

	"github.com/XoDeR/nethub-go/internal/db"
	"github.com/XoDeR/nethub-go/internal/env"
	"github.com/XoDeR/nethub-go/internal/store"
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
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres@localhost/nethub-go?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	// enable logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// test logger
	logger.Info("NetHub", zap.String("status", "running"))

	// db
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("database connected")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		logger: logger,
		store:  store,
	}

	// test logger, TODO: remove when not needed
	logger.Info(app)

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
