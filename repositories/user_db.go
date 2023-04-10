package repositories

import "github.com/jmoiron/sqlx"

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return userRepository{db: db}
}

func (r userRepository) GetUser(username string) (*User, error) {
	query := `
			SELECT username, password
			FROM users
			WHERE username = ?
			`
	user := User{}
	err := r.db.Get(&user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) CreateUser(user User) (*User, error) {
	query := `
			INSERT INTO users(username, password)
			VALUE (?, ?)
	`
	_, err := r.db.Exec(query, user.Username, user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
