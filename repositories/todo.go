package repositories

import "time"

type Todo struct {
	TodoID     string    `db:"todo_id"`
	Body       string    `db:"body"`
	Status     bool      `db:"status"`
	CreateDate time.Time `db:"create_date"`
	Username   string    `db:"username"`
}

type TodoRepository interface {
	GetAll(string) ([]Todo, error)
	GetById(string, string) (*Todo, error)
	CreateTodo(Todo) (*Todo, error)
	UpdateTodo(string, Todo) (*Todo, error)
	DeleteTodo(string, string) error
}
