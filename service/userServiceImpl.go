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

type UserServiceImpl struct {
	repo repository.UserRepository
	jwt *util.JwtUtil
}

func (r *UserServiceImpl) Login(loginRequest *dto.LoginRequest) (error, string) {
	err, user := r.repo.FindUserByEmail(loginRequest.Email)
	if err != nil {
		log.Println("No USER EXISTS, in adminService file")
		return errors.New("No User exists"), ""
	}

	if user.Password != loginRequest.Password {
		log.Println("Password Mismatch, in adminService file")
		return errors.New("Password Mismatch"), ""
	}

	if user.Role != "user" {
		log.Println("Unauthorized, in adminService file")
		return errors.New("Unauthorized access"), ""
	}

	err, token := r.jwt.CreateToken(loginRequest.Email, "user")
	if err != nil {
		return errors.New("Token NOT generated"), ""
	}
	return nil, token
}

func (us *UserServiceImpl) RegisterUser(user *entity.User) (error, *entity.User) {
	err, user := us.repo.RegisterUser(user)
	if err != nil {
		log.Println("User not added, adminService file")
		return err, user
	}
	return err, user
}

func NewUserService(repository repository.UserRepository, jwt *util.JwtUtil) service.UserService {
	return &UserServiceImpl{
		repo: repository,
		jwt: jwt,
	}
}