package repositories

import (
	"database/sql"
	"errors"
)

type userRepositoryMock struct {
	users []User
}

func NewUserRepositoryMock() UserRepository {
	users := []User{
		{Username: "jonh", Password: "$2a$04$e89IUjZbD13hwQ363gKUE.hXMbt6CnKsn7GkMY3NSK3H/xF/ET41C"},
		{Username: "bob", Password: "$2a$04$EM6aXTpGI0QqVEyoUYNZ1u2mvTHLghsl9moXT/tCax0FzjAfDaere"},
		{Username: "alan", Password: "$2a$04$eA3i9F1uOOGH2V78ycFgG.iS6jvBUAFoM1v/aFhH6xYVSbURTj10m"},
		{Username: "bas", Password: "$2a$04$XOB72JhV7Os8n7X2jjXZyezWPgnTbuYQIJZLxKjKgv8tb5qE1bnPC"},
	}
	return &userRepositoryMock{users: users}
}

func (r *userRepositoryMock) GetUser(username string) (*User, error) {
	for _, u := range r.users {
		if u.Username == username {
			return &u, nil
		}
	}
	return nil, sql.ErrNoRows
}

func (r *userRepositoryMock) CreateUser(user User) (*User, error) {
	for _, u := range r.users {
		if u.Username == user.Username {
			return nil, errors.New("username is duplicate")
		}
	}
	r.users = append(r.users, user)
	return &user, nil
}
