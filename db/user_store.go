package db

import (
	"context"
	"database/sql"
	"go/types"
	"time"
)

type UserStore struct {
	Db *sql.DB
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

func (s *UserStore) GetUserById(id int) (types.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user types.User
	statement := `select * from users where id = $1`

	err := s.Db.QueryRowContext(ctx, statement, id).Scan(&user)
	if err != nil {
		return types.User{}, nil
	}
	return user, nil
}
