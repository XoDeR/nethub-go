package store

import (
	"context"
	"database/sql"
)

type password struct {
	text *string
	hash []byte
}

type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  password `json:"-"`
	CreatedAt string   `json:"created_at"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, tx *sql.Tx, user *User) error {
	return nil
}

func (s *UserStore) GetByID(ctx context.Context, userID int64) (*User, error) {
	return nil, ErrNotFound
}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	return nil, ErrNotFound
}

func (s *UserStore) Delete(ctx context.Context, userID int64) error {
	return nil
}
