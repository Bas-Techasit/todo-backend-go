package repositories

import "github.com/jmoiron/sqlx"

type todoRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) TodoRepository {
	return todoRepositoryDB{db: db}
}

func (r todoRepositoryDB) GetAll() ([]Todo, error) {
	todos := []Todo{}
	query := "SELECT id, body, complete, create_date FROM todo"
	err := r.db.Select(&todos, query)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r todoRepositoryDB) GetById(id int) (*Todo, error) {
	todo := Todo{}
	query := `
				SELECT id, body, complete, create_date 
				FROM todo 
				WHERE id = ?
			`
	err := r.db.Get(&todo, query, id)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r todoRepositoryDB) CreateTodo(todo Todo) (*Todo, error) {
	query := `
				INSERT INTO todo(id, body, complete, create_date) 
				VALUE (?, ?, ?, ?)
			`
	_, err := r.db.Exec(
		query,
		todo.Id,
		todo.Body,
		todo.Complete,
		todo.CeateDate,
	)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}
