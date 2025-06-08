package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Users interface {
		Create(context.Context, *User) error
		GetByID(context.Context, int64) (*User, error)
		GetByEmail(context.Context, string) (*User, error)
		Delete(context.Context, int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UserStore{db},
	}
}

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
