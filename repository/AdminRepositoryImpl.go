package repository

import (
	"errors"
	"log"
	"raedmajeed/entity"
	"raedmajeed/repository/interfaces"

	"gorm.io/gorm"
)

type AdminRepositoryImpl struct {
	DB *gorm.DB
}


func (db *AdminRepositoryImpl) CreateUser(user *entity.User) (error, *entity.User) {
	if db.DB == nil {
			log.Println("Error connecting DB in FindUserById method, AdminRepositoryImpl package")
			return errors.New("Error Coonecting Database"), nil
	}

	err, user := db.FindUserByEmail(user.Email)
	if err != nil {
		log.Println("USER ALREADY EXISTS")
		return errors.New("User exists in db"), nil
	}

	result := db.DB.Create(&user)
	if result.Error != nil {
			log.Println("Unable to add user, AdminRepositoryImpl package")
			return errors.New("User not added to db"), user
	}

	return nil, user
}

func (db *AdminRepositoryImpl) DeleteUser(id int) (error, *entity.User) {
	if db.DB == nil {
		log.Println("Error connecting DB in FindUserById method, AdminRepositoryImpl package")
		return errors.New("Error Connecting Database"), nil
	}
	err, existingUser := db.FindUserById(id)
	if err != nil {
		log.Println("user not found, AdminRepositoryImpl package")
		return gorm.ErrRecordNotFound, nil
	}

	result := db.DB.Delete(existingUser)
	if result.Error != nil {
		log.Println("Unable to Delete User, AdminRepositoryImpl package")
		return errors.New("Error deleteing user"), existingUser
	}

	return nil, existingUser
}

func (db *AdminRepositoryImpl) FindUserById(id int) (error, *entity.User) {
	if db.DB == nil {
		log.Println("Error connecting DB in FindUserById method, AdminRepositoryImpl package")
		return errors.New("Error Connecting Database"), nil
	}

	user := &entity.User{}
	result := db.DB.First(user, id)

	if result.Error != nil {
		log.Println("ID not found in DB")
		return errors.New("Id not Found in DB"), nil
	}

	return nil, user
}

func (db *AdminRepositoryImpl) FindUserByEmail(email string) (error, *entity.User) {
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


func (db *AdminRepositoryImpl) FindUserContaining(str string) (error, []*entity.User) {
	if db.DB == nil {
		log.Println("Error connecting DB in FindUserById method, AdminRepositoryImpl package")
		return errors.New("Error Connecting Database"), nil
	}

	users := []*entity.User{}
	result := db.DB.Where("email LIKE ?", "%"+str+"%").Find(&users)
	if result.Error != nil {
		return result.Error, nil
	}
	
	return nil, users
}


func (db *AdminRepositoryImpl) UpdateUser(id int, user *entity.User) (error, *entity.User) {
	if db.DB == nil {
		log.Println("DB failes in UpdateUser, AdminRepositoryImpl package")
		return errors.New("DB connection Failed"), nil
	}

	err, existingUser := db.FindUserById(id)
	if err != nil {
		log.Println("user not found, AdminRepositoryImpl package")
		return gorm.ErrRecordNotFound, nil
	}

	if user.Email != "" {existingUser.Email = user.Email}
	if user.PhoneNumber != "" {existingUser.PhoneNumber = user.PhoneNumber}
	if user.Password != "" {existingUser.Password = user.Password}

	result := db.DB.Save(existingUser)
	if result.Error != nil {
		log.Println("User Not Updated maybe the same email already present, AdminRepositoryImpl package")
		return errors.New("User not updated"), nil
	}

	return nil, existingUser
}

func (db *AdminRepositoryImpl) FindAllusers() (error, []*entity.User) {
	if db.DB == nil {
		log.Println("DB failes in UpdateUser, AdminRepositoryImpl package")
		return errors.New("DB connection Failed"), nil
	}
	userList := []*entity.User{}
	db.DB.Find(&userList)
	return nil, userList
}

func (db *AdminRepositoryImpl) BlockUser(id int) (error, *entity.User) {
	if db.DB == nil {
		log.Println("DB failes in UpdateUser, AdminRepositoryImpl package")
		return errors.New("DB connection Failed"), nil
	}
	err, user := db.FindUserById(id)
    if err != nil {
        log.Println("User not found, AdminRepositoryImpl package")
        return gorm.ErrRecordNotFound, nil
    }

    user.IsAccountLocked = true

    result := db.DB.Save(user)
    if result.Error != nil {
				log.Println("Unable to block user")
        return result.Error, nil
    }

    return nil, user
}

func (db *AdminRepositoryImpl) UnBlockUser(id int) (error, *entity.User) {
	if db.DB == nil {
		log.Println("DB failes in UpdateUser, AdminRepositoryImpl package")
		return errors.New("DB connection Failed"), nil
	}
	err, user := db.FindUserById(id)
    if err != nil {
        log.Println("User not found, AdminRepositoryImpl package")
        return gorm.ErrRecordNotFound, nil
    }

    user.IsAccountLocked = false

    result := db.DB.Save(user)
    if result.Error != nil {
				log.Println("Unable to unblock user")
        return result.Error, nil
    }

    return nil, user
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepository {
	return &AdminRepositoryImpl{
		DB: db,
	}
}
