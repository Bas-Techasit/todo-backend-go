package services

type TodoService interface {
	GetTodos() ([]TodoResponse, error)
	GetTodo(int) (*TodoResponse, error)
	NewTodo(NewTodoRequest) (*TodoResponse, error)
	EditTodo(int, EditTodoRequest) (*TodoResponse, error)
	DeleteTodo(int) error
}

type NewTodoRequest struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
}

type EditTodoRequest struct {
	Body     string `json:"body"`
	Complete bool   `json:"complete"`
}

type TodoResponse struct {
	Id       int    `json:"id"`
	Body     string `json:"body"`
	Complete bool   `json:"complete"`
}
