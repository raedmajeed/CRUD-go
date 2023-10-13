package interfaces

import (
	"raedmajeed/dto"
	"raedmajeed/entity"
)

type AdminService interface {
	Login(loginRequest *dto.LoginRequest) (error, string)
	DeleteUser(id int) (error, *entity.User)
	UpdateUser(id int, user entity.User) (error, *entity.User)
	BlockUser(id int) (error, *entity.User)
	UnBlockUser(id int) (error, *entity.User)
	AddUser(user *entity.User) (error, *entity.User)
	FindUser(id int) (error, *entity.User)
	SearchUser(str string) (error, []*entity.User)
	FindAllUsers() (error, []*entity.User)
}