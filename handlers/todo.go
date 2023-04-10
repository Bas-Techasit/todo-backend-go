package handlers

import (
	"todo-backend/logs"
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
	username := c.Params("username")
	todos, err := h.todoSrv.GetTodos(username)
	if err != nil {
		logs.Error(err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	return c.JSON(todos)
}

// func (h todoHandler) GetTodo(c *fiber.Ctx) error {
// 	todoId := c.Params("todoId")

// 	todo, err := h.todoSrv.GetTodo(todoId)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusNotFound, "todo not found")
// 	}

// 	return c.JSON(todo)
// }

func (h todoHandler) NewTodo(c *fiber.Ctx) error {
	username := c.Params("username")
	todo := services.NewTodoRequest{}
	err := c.BodyParser(&todo)
	if err != nil {
		return fiber.ErrUnprocessableEntity
	}

	todoResponse, err := h.todoSrv.NewTodo(username, todo)
	if err != nil {
		return err
	}
	c.SendStatus(fiber.StatusCreated)
	return c.JSON(todoResponse)
}

// func (h todoHandler) UpdateTodo(c *fiber.Ctx) error {
// 	todoId := c.Params("todoId")

// 	todo := services.EditTodoRequest{}
// 	err := c.BodyParser(&todo)
// 	if err != nil {
// 		return fiber.ErrUnprocessableEntity
// 	}

// 	todoResponse, err := h.todoSrv.EditTodo(todoId, todo)
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(todoResponse)
// }

// func (h todoHandler) DeleteTodo(c *fiber.Ctx) error {
// 	todoId := c.Params("todoId")
// 	return h.todoSrv.DeleteTodo(todoId)
// }
