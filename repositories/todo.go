package repositories

import "time"

type Todo struct {
	Id         string    `db:"id"`
	Body       string    `db:"body"`
	Complete   bool      `db:"complete"`
	CreateDate time.Time `db:"create_date"`
}

type TodoRepository interface {
	GetAll() ([]Todo, error)
	GetById(string) (*Todo, error)
	CreateTodo(Todo) (*Todo, error)
	UpdateTodo(string, string, bool) (*Todo, error)
	DeleteTodo(string) error
}
