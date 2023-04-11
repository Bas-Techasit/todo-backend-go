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
			TodoID: t.TodoID,
			Body:   t.Body,
			Status: t.Status,
		}
		todoResponses = append(todoResponses, todoResponse)
	}
	return todoResponses, nil
}

func (s todoService) NewTodo(username string, res NewTodoRequest) (*TodoResponse, error) {

	if res.Body == "" {
		return nil, fiber.NewError(
			fiber.StatusUnprocessableEntity,
			"body not empty",
		)
	}

	todo := repositories.Todo{
		TodoID:     uuid.NewString(),
		Body:       res.Body,
		Status:     false,
		CreateDate: time.Now(),
		Username:   username,
	}

	created, err := s.todoRepo.CreateTodo(todo)
	if err != nil {
		logs.Error(err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "unexpected error")
	}

	todoResponse := TodoResponse{
		TodoID: created.TodoID,
		Body:   created.Body,
		Status: created.Status,
	}
	return &todoResponse, nil
}

func (s todoService) EditTodo(username, todoID string, updateReq UpdateTodoRequest) (*TodoResponse, error) {
	oldTodo, err := s.todoRepo.GetById(username, todoID)
	if err != nil {
		return nil, fiber.NewError(
			fiber.StatusNotFound,
			"todo not found",
		)
	}

	if updateReq.Body == "" {
		updateReq.Body = oldTodo.Body
	}

	todo := repositories.Todo{
		TodoID:     todoID,
		Body:       updateReq.Body,
		Status:     updateReq.Status,
		CreateDate: oldTodo.CreateDate,
		Username:   username,
	}

	var updatedTodo *repositories.Todo
	updatedTodo, err = s.todoRepo.UpdateTodo(username, todo)
	if err != nil {
		logs.Error(err)
		return nil, fiber.ErrInternalServerError
	}

	todoResponse := TodoResponse{
		TodoID: updatedTodo.TodoID,
		Body:   updatedTodo.Body,
		Status: updatedTodo.Status,
	}

	return &todoResponse, nil
}

func (s todoService) DeleteTodo(username, todoID string) error {
	err := s.todoRepo.DeleteTodo(username, todoID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.NewError(fiber.StatusNotFound, "todo not found")
		}
		logs.Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, "unexpected error")
	}

	return nil
}
