package repository

import (
	"fmt"
	"raedmajeed/repository/interfaces"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (ur *UserRepositoryImpl) CreateUser() error {
	fmt.Println(ur.db)
	return nil
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &UserRepositoryImpl {
		db: db,
	}
}