package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/IDOMATH/auth/types"
	"golang.org/x/crypto/bcrypt"
)

type UserStore struct {
	Db *sql.DB
}

func NewUserStore(conn *sql.DB) *UserStore {
	return &UserStore{
		Db: conn,
	}
}

func (s *UserStore) InsertUser(user types.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newId int
	statement := `insert into users (username, email, password_hash)
				  values ($1, $2, $3) returning id`

	err := s.Db.QueryRowContext(ctx, statement, user.Username, user.Email, user.PasswordHash).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (s *UserStore) Authenticate(username, password string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	var passwordHash string
	statement := `select id, password_hash from users where username = $1`

	err := s.Db.QueryRowContext(ctx, statement, username).Scan(&id, passwordHash)
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return 0, err
	}
	return id, nil
}
