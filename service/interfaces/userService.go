package interfaces

import (
	"raedmajeed/dto"
	"raedmajeed/entity"
)

type UserService interface {
	Login(loginRequest *dto.LoginRequest) (error, string)
	RegisterUser(user *entity.User) (error, *entity.User)
}