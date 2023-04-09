package repositories

import "time"

type Todo struct {
	Id         int       `db:"id"`
	Body       string    `db:"body"`
	Complete   bool      `db:"complete"`
	CreateDate time.Time `db:"create_date"`
}

type TodoRepository interface {
	GetAll() ([]Todo, error)
	GetById(int) (*Todo, error)
	CreateTodo(Todo) (*Todo, error)
	UpdateTodo(int, string, bool) (*Todo, error)
	DeleteTodo(int) error
}
