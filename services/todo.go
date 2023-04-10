package services

type TodoService interface {
	GetTodos(string) ([]TodoResponse, error)
	// GetTodo(string) (*TodoResponse, error)
	NewTodo(string, NewTodoRequest) (*TodoResponse, error)
	// EditTodo(string, EditTodoRequest) (*TodoResponse, error)
	DeleteTodo(string) error
}

type NewTodoRequest struct {
	Body string `json:"body"`
}

type EditTodoRequest struct {
	Body     string `json:"body"`
	Complete bool   `json:"complete"`
}

type TodoResponse struct {
	Id       string `json:"id"`
	Body     string `json:"body"`
	Complete bool   `json:"complete"`
}
