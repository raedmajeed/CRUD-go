package interfaces

import "raedmajeed/entity"

type UserRepository interface {
	RegisterUser(user *entity.User) (error, *entity.User)
	FindUserByEmail(email string) (error, *entity.User)
}