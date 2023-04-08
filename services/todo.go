package services

type TodoService interface {
	GetTodos() ([]TodoResponse, error)
	GetTodo(int) (*TodoResponse, error)
}

type TodoResponse struct {
	Id       int    `json:"id"`
	Body     string `json:"body"`
	Complete bool   `json:"complete"`
}
