package repository

import (
	"raedmajeed/repository/interfaces"

	"gorm.io/gorm"
)

type AdminRepositoryImpl struct {
}

func (*AdminRepositoryImpl) CreateUser() error {
	panic("unimplemented")
}

func (*AdminRepositoryImpl) DeleteUser() {
	panic("unimplemented")
}

func (*AdminRepositoryImpl) FindUserById() {
	panic("unimplemented")
}


func (*AdminRepositoryImpl) FindUserContaining() {
	panic("unimplemented")
}


func (*AdminRepositoryImpl) UpdateUser() {
	panic("unimplemented")
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepository {
	return &AdminRepositoryImpl{}
}
