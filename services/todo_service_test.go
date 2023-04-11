package services_test

import (
	"reflect"
	"testing"
	"todo-backend/repositories"
	"todo-backend/services"

	"github.com/gofiber/fiber/v2"
)

var todoRepoMock = repositories.NewTodoRepositoryMock()
var todoService = services.NewTodoService(todoRepoMock)

func TestGetTodos(t *testing.T) {
	t.Run("get todos should be pass", func(t *testing.T) {
		got, err := todoService.GetTodos("techasit")
		want := []services.TodoResponse{
			{TodoID: "1", Body: "item 1", Status: false},
			{TodoID: "2", Body: "item 2", Status: false},
			{TodoID: "3", Body: "item 3", Status: false},
			{TodoID: "4", Body: "item 4", Status: false},
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("get todos should not pass", func(t *testing.T) {
		_, err := todoService.GetTodos("xxxxx")
		want := fiber.NewError(fiber.StatusInternalServerError, "unexpected error")

		if err == nil {
			t.Errorf("don't have a error but i want")
			return
		}

		if !reflect.DeepEqual(err, want) {
			t.Errorf("got %v want %v", err, want)
		}
	})

}

func TestNewTodo(t *testing.T) {

	t.Run("New a todo should be pass", func(t *testing.T) {
		newTodo := services.NewTodoRequest{
			Body: "Test new todo",
		}
		got, err := todoService.NewTodo("techasit", newTodo)
		want := "Test new todo"
		if err != nil {
			t.Errorf(err.Error())
		}
		if got.Body != want {
			t.Errorf("got %v want %v", got.Body, want)
		}
	})

	t.Run("New a todo with empty body should not pass", func(t *testing.T) {
		newTodo := services.NewTodoRequest{
			Body: "",
		}
		_, err := todoService.NewTodo("techasit", newTodo)
		want := fiber.NewError(
			fiber.StatusUnprocessableEntity,
			"body not empty",
		)
		if err == nil {
			t.Errorf("don't have a error but i want")
			return
		}

		if !reflect.DeepEqual(err, want) {
			t.Errorf("got %v want %v", err, want)
		}
	})

}

func TestEditTodo(t *testing.T) {

	t.Run("user techasit updates body id = 1 should be pass", func(t *testing.T) {
		todo := services.UpdateTodoRequest{
			Body: "Test update todo",
		}
		got, err := todoService.EditTodo("techasit", "1", todo)
		want := &services.TodoResponse{
			TodoID: "1",
			Body:   "Test update todo",
			Status: false,
		}
		if err != nil {
			t.Errorf(err.Error())
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got.Body, want)
		}
	})

	t.Run("user techasit updates todo that not exist id should not pass", func(t *testing.T) {
		todo := services.UpdateTodoRequest{
			Body: "Test update todo",
		}
		_, err := todoService.EditTodo("techasit", "5", todo)
		want := fiber.NewError(
			fiber.StatusNotFound,
			"todo not found",
		)
		if err == nil {
			t.Errorf("don't have a error but i want")
			return
		}

		if !reflect.DeepEqual(err, want) {
			t.Errorf("got %v want %v", err, want)
		}
	})

}

func TestDeleteTodo(t *testing.T) {
	t.Run("user techasit delete id = 3 should be pass", func(t *testing.T) {
		err := todoService.DeleteTodo("techasit", "3")
		if err != nil {
			t.Errorf("got %v want %v", err, nil)
		}
	})

	t.Run("user techasit delete todo id = 5 should not pass", func(t *testing.T) {
		err := todoService.DeleteTodo("techasit", "5")
		want := fiber.NewError(fiber.StatusNotFound, "todo not found")
		if err == nil {
			t.Errorf("don't have a error but i want")
			return
		}

		if !reflect.DeepEqual(err, want) {
			t.Errorf("got %v want %v", err, want)
		}
	})
}
