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

func (r todoRepositoryDB) GetById(username, todoID string) (*Todo, error) {
	todo := Todo{}
	query := `
			SELECT todo_id, body, status, create_date, username
			FROM todos
			WHERE todo_id = ? AND username = ?
	`
	err := r.db.Get(&todo, query, todoID, username)
	if err != nil {
		return nil, err
	}
	return &todo, nil
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

func (r todoRepositoryDB) UpdateTodo(username string, todo Todo) (*Todo, error) {
	query := `
				UPDATE todos
				SET todo_id = ?, body = ?, status = ?, create_date = ?, username = ?
				WHERE todo_id = ?
			`
	_, err := r.db.Exec(
		query,
		todo.TodoID,
		todo.Body,
		todo.Status,
		todo.CreateDate,
		todo.Username,
		todo.TodoID,
	)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r todoRepositoryDB) DeleteTodo(username, todoID string) error {
	query := `
				DELETE FROM todos
				WHERE todo_id = ? AND username = ?
			`
	_, err := r.db.Exec(query, todoID, username)
	if err != nil {
		return err
	}

	return nil
}
