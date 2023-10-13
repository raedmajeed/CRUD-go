package service

import (
	"errors"
	"log"
	"raedmajeed/dto"
	"raedmajeed/entity"
	repository "raedmajeed/repository/interfaces"
	service "raedmajeed/service/interfaces"
	"raedmajeed/util"
)

type AdminServiceImpl struct {
	repo repository.AdminRepository
	jwt *util.JwtUtil
}

func (r *AdminServiceImpl) AddUser(user *entity.User) (error, *entity.User) {
	err, user := r.repo.CreateUser(user)
	if err != nil {
		log.Panicln("User not added, adminService file")
		return err, user
	}
	return err, user
}

func (r *AdminServiceImpl) BlockUser(id int) (error, *entity.User) {
	err, user := r.repo.BlockUser(id)
	if err != nil {
		log.Println("Error Blocking user, in adminServiceImpl file")
		return err, user
	}
	return err, user
}

func (r *AdminServiceImpl) UnBlockUser(id int) (error, *entity.User) {
	err, user := r.repo.UnBlockUser(id)
	if err != nil {
		log.Println("Error Blocking user, in adminServiceImpl file")
		return err, user
	}
	return err, user
}

func (r *AdminServiceImpl) DeleteUser(id int) (error, *entity.User) {
	err, user := r.repo.DeleteUser(id)
	if err != nil {
		log.Println("Error Deleting user, in adminServiceImpl file")
		return err, user
	}
	return err, user
}


func (r *AdminServiceImpl) FindAllUsers() (error, []*entity.User) {
	err, userList := r.repo.FindAllusers()
	if err != nil {
		log.Println("Error Finding All users, in adminServiceImpl file")
		return err, userList
	}
	return err, userList
}

func (r *AdminServiceImpl) FindUser(id int) (error, *entity.User) {
	err, user := r.repo.FindUserById(id)
	if err != nil {
		log.Println("Error Finding user, in adminServiceImpl file")
		return err, user
	}
	
	return err, user
}


func (r *AdminServiceImpl) Login(loginRequest *dto.LoginRequest) (error, string) {
	err, user := r.repo.FindUserByEmail(loginRequest.Email)
	if err != nil {
		log.Println("No USER EXISTS, in adminService file")
		return errors.New("No User exists"), ""
	}

	if user.Password != loginRequest.Password{
		log.Println("Password Mismatch, in adminService file")
		return errors.New("Password Mismatch"), ""
	}

	if user.Role != "admin" {
		log.Println("Unauthorized, in adminService file")
		return errors.New("Unauthorized access"), ""
	}

	err, token := r.jwt.CreateToken(loginRequest.Email, "admin")
	if err != nil {
		return errors.New("Token NOT generated"), ""
	}

	return nil, token
}


func (r *AdminServiceImpl) SearchUser(str string) (error, []*entity.User) {
	err, userList := r.repo.FindUserContaining(str)
	if err != nil {
		log.Println("Error Searching users, in adminServiceImpl file")
		return err, userList
	}
	
	return err, userList
}


func (r *AdminServiceImpl) UpdateUser(id int, user entity.User) (error, *entity.User) {
	err, updatedUser := r.repo.UpdateUser(id, &user)

	if err != nil {
		log.Println("Error Updating user, in adminServiceImpl file")
		return err, updatedUser
	}

	return nil, updatedUser
}

func NewAdminService(repository repository.AdminRepository, jwt *util.JwtUtil) service.AdminService {
	return &AdminServiceImpl{
		repo: repository,
		jwt: jwt,
	}
}
