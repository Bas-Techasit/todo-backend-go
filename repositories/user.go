package repositories

type User struct {
	Username string `db:"username"`
	Password string `db:"password"`
}

type UserRepository interface {
	GetUser(string) (*User, error)
	CreateUser(string, string) (*User, error)
}
