package repositories

type Todo struct {
	Id        int    `db:"id"`
	Body      string `db:"body"`
	Complete  bool   `db:"complete"`
	CeateDate string `db:"create_date"`
}

type TodoRepository interface {
	GetAll() ([]Todo, error)
	GetById(int) (*Todo, error)
}
