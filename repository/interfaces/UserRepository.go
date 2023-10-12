package interfaces

type UserRepository interface {
	CreateUser() error
}