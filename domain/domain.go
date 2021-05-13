package domain

type UserRepo interface {
	GetUsers()
}

type User struct {
	ID   int
	Name string
}
