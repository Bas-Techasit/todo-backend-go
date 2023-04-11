package repositories

import (
	"errors"
	"time"
)

type todoRepositoryMock struct {
	todos []Todo
}

func NewTodoRepositoryMock() TodoRepository {
	todos := []Todo{
		{TodoID: "1", Body: "item 1", Status: false, CreateDate: time.Now(), Username: "techasit"},
		{TodoID: "2", Body: "item 2", Status: false, CreateDate: time.Now(), Username: "techasit"},
		{TodoID: "3", Body: "item 3", Status: false, CreateDate: time.Now(), Username: "techasit"},
		{TodoID: "4", Body: "item 4", Status: false, CreateDate: time.Now(), Username: "techasit"},
		{TodoID: "5", Body: "item 1", Status: false, CreateDate: time.Now(), Username: "bob"},
		{TodoID: "6", Body: "item 2", Status: false, CreateDate: time.Now(), Username: "bob"},
		{TodoID: "7", Body: "item 3", Status: false, CreateDate: time.Now(), Username: "bob"},
		{TodoID: "8", Body: "item 4", Status: false, CreateDate: time.Now(), Username: "bob"},
	}
	return &todoRepositoryMock{todos: todos}
}

func (r *todoRepositoryMock) GetAll(username string) ([]Todo, error) {
	res := []Todo{}
	for _, t := range r.todos {
		if t.Username == username {
			res = append(res, t)
			return res, nil
		}
	}
	return nil, errors.New("username not found")
}

func (r *todoRepositoryMock) GetById(username, todoID string) (*Todo, error) {
	for _, t := range r.todos {
		if t.Username == username && t.TodoID == todoID {
			return &t, nil
		}
	}
	return nil, errors.New("todo not found")
}

func (r *todoRepositoryMock) CreateTodo(todo Todo) (*Todo, error) {
	r.todos = append(r.todos, todo)
	return &todo, nil
}

func (r *todoRepositoryMock) UpdateTodo(username string, todo Todo) (*Todo, error) {
	for index, t := range r.todos {
		if t.Username == username && t.TodoID == todo.TodoID {
			r.todos[index] = todo
			return &r.todos[index], nil
		}
	}
	return nil, errors.New("todo not found")
}

func (r *todoRepositoryMock) DeleteTodo(username, todoID string) error {
	for index, t := range r.todos {
		if t.Username == username && t.TodoID == todoID {
			r.todos = append(r.todos[:index], r.todos[index+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}
