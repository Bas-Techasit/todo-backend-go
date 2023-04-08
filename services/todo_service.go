package services

import (
	"database/sql"
	"todo-backend/errs"
	"todo-backend/logs"
	"todo-backend/repositories"
)

type todoService struct {
	todoRepo repositories.TodoRepository
}

func NewTodoService(todoRepo repositories.TodoRepository) TodoService {
	return todoService{todoRepo: todoRepo}
}

func (s todoService) GetTodos() ([]TodoResponse, error) {
	todos, err := s.todoRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	todoResponses := []TodoResponse{}
	for _, t := range todos {
		todoResponse := TodoResponse{
			Id:       t.Id,
			Body:     t.Body,
			Complete: t.Complete,
		}
		todoResponses = append(todoResponses, todoResponse)
	}
	return todoResponses, nil
}

func (s todoService) GetTodo(id int) (*TodoResponse, error) {
	todo, err := s.todoRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("todo not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()

	}

	todoResponse := TodoResponse{
		Id:       todo.Id,
		Body:     todo.Body,
		Complete: todo.Complete,
	}
	return &todoResponse, nil
}
