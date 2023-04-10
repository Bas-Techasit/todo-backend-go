package repositories

type Todo struct {
	TodoID     string `db:"todo_id"`
	Body       string `db:"body"`
	Status     bool   `db:"status"`
	CreateDate string `db:"create_date"`
	Username   string `db:"username"`
}

type TodoRepository interface {
	GetAll(string) ([]Todo, error)
	CreateTodo(Todo) (*Todo, error)
	// UpdateTodo(string, string, bool) (*Todo, error)
	DeleteTodo(string) error
}
