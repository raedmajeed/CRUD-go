package interfaces

type AdminService interface {
	Login()
	DeleteUser()
	UpdateUser()
	BlockUser()
	AddUser()
	FindUser()
	SearchUser()
	FindAllUsers()
}