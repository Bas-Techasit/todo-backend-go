package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-backend/services"

	"github.com/gorilla/mux"
)

type todoHandler struct {
	todoSrv services.TodoService
}

func NewTodoHandler(todoSrv services.TodoService) todoHandler {
	return todoHandler{todoSrv: todoSrv}
}

func (h todoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoSrv.GetTodos()
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (h todoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["todoId"])
	todo, err := h.todoSrv.GetTodo(id)
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
