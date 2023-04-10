package services

import (
	"database/sql"
	"time"
	"todo-backend/logs"
	"todo-backend/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type todoService struct {
	todoRepo repositories.TodoRepository
}

func NewTodoService(todoRepo repositories.TodoRepository) TodoService {
	return todoService{todoRepo: todoRepo}
}

func (s todoService) GetTodos(username string) ([]TodoResponse, error) {
	todos, err := s.todoRepo.GetAll(username)
	if err != nil {
		logs.Error(err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "unexpected error")
	}
	todoResponses := []TodoResponse{}
	for _, t := range todos {
		todoResponse := TodoResponse{
			Id:       t.TodoID,
			Body:     t.Body,
			Complete: t.Status,
		}
		todoResponses = append(todoResponses, todoResponse)
	}
	return todoResponses, nil
}

// func (s todoService) GetTodo(id string) (*TodoResponse, error) {
// 	todo, err := s.todoRepo.GetById(id)
// 	if err != nil {

// 		if err == sql.ErrNoRows {
// 			return nil, fiber.NewError(fiber.StatusNotFound, "todo not fonud")
// 		}

// 		logs.Error(err)
// 		return nil, fiber.NewError(fiber.StatusInternalServerError, "unexpected error")

// 	}

// 	todoResponse := TodoResponse{
// 		Id:       todo.Id,
// 		Body:     todo.Body,
// 		Complete: todo.Complete,
// 	}
// 	return &todoResponse, nil
// }

func (s todoService) NewTodo(username string, res NewTodoRequest) (*TodoResponse, error) {

	if res.Body == "" {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "body not empty")
	}

	todo := repositories.Todo{
		TodoID:     uuid.NewString(),
		Body:       res.Body,
		Status:     false,
		CreateDate: time.Now().Format("2006-1-2 15:04:05"),
		Username:   username,
	}

	created, err := s.todoRepo.CreateTodo(todo)
	if err != nil {
		logs.Error(err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "unexpected error")
	}

	todoResponse := TodoResponse{
		Id:       created.TodoID,
		Body:     created.Body,
		Complete: created.Status,
	}
	return &todoResponse, nil
}

// func (s todoService) EditTodo(id string, e EditTodoRequest) (*TodoResponse, error) {
// 	if e.Body == "" {
// 		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "body not empty")
// 	}

// 	updatedTodo, err := s.todoRepo.UpdateTodo(id, e.Body, e.Complete)
// 	if err != nil {
// 		logs.Error(err)
// 		return nil, fiber.NewError(fiber.StatusInternalServerError, "unexpected error")
// 	}

// 	todoResponse := TodoResponse{
// 		Id:       updatedTodo.Id,
// 		Body:     updatedTodo.Body,
// 		Complete: updatedTodo.Complete,
// 	}

// 	return &todoResponse, nil
// }

func (s todoService) DeleteTodo(id string) error {
	err := s.todoRepo.DeleteTodo(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "todo not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "unexpected error")
	}

	return nil
}
