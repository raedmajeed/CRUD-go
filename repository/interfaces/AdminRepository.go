package interfaces

import "raedmajeed/entity"

type AdminRepository interface {
	CreateUser(user *entity.User) (error, *entity.User)
	FindUserById(id int) (error, *entity.User)
	DeleteUser(id int) (error, *entity.User)
	UpdateUser(id int, user *entity.User) (error, *entity.User)
	FindUserContaining(str string) (error, []*entity.User)
	FindUserByEmail(email string) (error, *entity.User)
	FindAllusers() (error, []*entity.User)
	BlockUser(id int) (error, *entity.User)
	UnBlockUser(id int) (error, *entity.User)
}