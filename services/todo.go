package services

type TodoService interface {
	GetTodos(string) ([]TodoResponse, error)
	NewTodo(string, NewTodoRequest) (*TodoResponse, error)
	EditTodo(string, string, UpdateTodoRequest) (*TodoResponse, error)
	DeleteTodo(string, string) error
}

type NewTodoRequest struct {
	Body string `json:"body"`
}

type UpdateTodoRequest struct {
	Body   string `json:"body"`
	Status bool   `json:"status"`
}

type TodoResponse struct {
	TodoID string `json:"id"`
	Body   string `json:"body"`
	Status bool   `json:"status"`
}
