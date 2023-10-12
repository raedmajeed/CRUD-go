package interfaces

type AdminRepository interface {
	CreateUser() error
	FindUserById()
	DeleteUser()
	UpdateUser()
	FindUserContaining()
}