package repositories

import "github.com/jmoiron/sqlx"

type todoRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) TodoRepository {
	return todoRepositoryDB{db: db}
}

func (r todoRepositoryDB) GetAll(username string) ([]Todo, error) {
	todos := []Todo{}
	query := `
				SELECT todo_id, body, status, create_date, username
				FROM todos
				WHERE username = ?
			`
	err := r.db.Select(&todos, query, username)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r todoRepositoryDB) CreateTodo(todo Todo) (*Todo, error) {
	query := `
				INSERT INTO todos(todo_id, body, status, create_date, username) 
				VALUE (?, ?, ?, ?, ?)
			`
	_, err := r.db.Exec(
		query,
		todo.TodoID,
		todo.Body,
		todo.Status,
		todo.CreateDate,
		todo.Username,
	)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

// func (r todoRepositoryDB) UpdateTodo(id string, body string, isCompleted bool) (*Todo, error) {
// 	query := `
// 				UPDATE todo
// 				SET body = ?, complete = ?
// 				WHERE id = ?
// 			`
// 	_, err := r.db.Exec(query, body, isCompleted, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	todo := Todo{
// 		Id:       id,
// 		Body:     body,
// 		Complete: isCompleted,
// 	}

// 	return &todo, nil
// }

func (r todoRepositoryDB) DeleteTodo(todoID string) error {
	query := `
				DELETE FROM todos
				WHERE id = ?
			`
	_, err := r.db.Exec(query, todoID)
	if err != nil {
		return err
	}

	return nil
}
