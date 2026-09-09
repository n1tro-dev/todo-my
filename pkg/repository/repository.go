package repository

type Authorizaton interface {

}

type TodoList interface {

}

type TodoItem interface {

}

type Repository struct {
	Authorizaton
	TodoList
	TodoItem
}

func NewRepository() *Repository{
	return &Repository{}
}