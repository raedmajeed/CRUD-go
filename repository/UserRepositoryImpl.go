package repository

import (
	"errors"
	"log"
	"raedmajeed/entity"
	"raedmajeed/repository/interfaces"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (db *UserRepositoryImpl) RegisterUser(user *entity.User) (error, *entity.User) {
	if db.DB == nil {
			log.Println("Error connecting DB in FindUserById method, AdminRepositoryImpl package")
			return errors.New("Error Coonecting Database"), nil
	}

	result := db.DB.Create(&user)
	if result.Error != nil {
			log.Println("Unable to add user, AdminRepositoryImpl package")
			return errors.New("User not added to db"), user
	}

	return nil, user
}

func (db *UserRepositoryImpl) FindUserByEmail(email string) (error, *entity.User) {
	if db.DB == nil {
		log.Println("Error connecting DB in FindUserById method, AdminRepositoryImpl package")
		return errors.New("Error Connecting Database"), nil
	}

	user := &entity.User{}
	result := db.DB.Where("email = ?", email).First(user)

	if result.Error != nil {
		log.Println("User not found in DB")
		return errors.New("User not Found in DB"), nil
	}

	return nil, user
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &UserRepositoryImpl {
		DB: db,
	}
}