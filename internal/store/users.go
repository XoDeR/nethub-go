package store

import (
	"context"
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail    = errors.New("a user with that email already exists")
	ErrDuplicateUsername = errors.New("a user with that username already exists")
)

type password struct {
	text *string
	hash []byte
}

func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.text = &text
	p.hash = hash

	return nil
}

func (p *password) Compare(text string) error {
	return bcrypt.CompareHashAndPassword(p.hash, []byte(text))
}

type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  password `json:"-"`
	CreatedAt string   `json:"created_at"`
	RoleID    int64    `json:"role_id"`
	Role      Role     `json:"role"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	return withTx(s.db, ctx, func(tx *sql.Tx) error {
		query := `
		INSERT INTO users (username, password, email, role_id) VALUES 
    ($1, $2, $3, (SELECT id FROM roles WHERE name = $4))
    RETURNING id, created_at
	`

		ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
		defer cancel()

		role := user.Role.Name
		if role == "" {
			role = "user"
		}

		err := tx.QueryRowContext(
			ctx,
			query,
			user.Username,
			user.Password.hash,
			user.Email,
			role,
		).Scan(
			&user.ID,
			&user.CreatedAt,
		)
		if err != nil {
			switch {
			case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
				return ErrDuplicateEmail
			case err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"`:
				return ErrDuplicateUsername
			default:
				return err
			}
		}

		return nil
	})
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
