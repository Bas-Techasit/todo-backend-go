package handlers

import (
	"todo-backend/services"

	"github.com/gofiber/fiber/v2"
)

type todoHandler struct {
	todoSrv services.TodoService
}

func NewTodoHandler(todoSrv services.TodoService) todoHandler {
	return todoHandler{todoSrv: todoSrv}
}

func (h todoHandler) GetTodos(c *fiber.Ctx) error {
	todos, err := h.todoSrv.GetTodos()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	return c.JSON(todos)
}

func (h todoHandler) GetTodo(c *fiber.Ctx) error {
	todoId, err := c.ParamsInt("todoId")
	if err != nil {
		return fiber.ErrBadRequest
	}

	todo, err := h.todoSrv.GetTodo(todoId)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "todo not found")
	}

	return c.JSON(todo)
}

func (h todoHandler) NewTodo(c *fiber.Ctx) error {
	todo := services.NewTodoRequest{}
	err := c.BodyParser(&todo)
	if err != nil {
		return fiber.ErrUnprocessableEntity
	}

	todoResponse, err := h.todoSrv.NewTodo(todo)
	if err != nil {
		return err
	}
	c.SendStatus(fiber.StatusCreated)
	return c.JSON(todoResponse)
}

func (h todoHandler) UpdateTodo(c *fiber.Ctx) error {
	todoId, err := c.ParamsInt("todoId")
	if err != nil {
		return fiber.ErrBadRequest
	}

	todo := services.EditTodoRequest{}
	err = c.BodyParser(&todo)

	if err != nil {
		return fiber.ErrUnprocessableEntity
	}

	todoResponse, err := h.todoSrv.EditTodo(todoId, todo)
	if err != nil {
		return err
	}

	return c.JSON(todoResponse)
}

func (h todoHandler) DeleteTodo(c *fiber.Ctx) error {
	todoId, err := c.ParamsInt("todoId")
	if err != nil {
		return fiber.ErrBadRequest
	}

	return h.todoSrv.DeleteTodo(todoId)

}
