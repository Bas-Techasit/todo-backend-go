package services

type TodoService interface {
	GetTodos() ([]TodoResponse, error)
	GetTodo(string) (*TodoResponse, error)
	NewTodo(NewTodoRequest) (*TodoResponse, error)
	EditTodo(string, EditTodoRequest) (*TodoResponse, error)
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
