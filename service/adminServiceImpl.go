package service

import (
	repository "raedmajeed/repository/interfaces"
	service "raedmajeed/service/interfaces"
)



type AdminServiceImpl struct {
}

// AddUser implements interfaces.AdminService.
func (*AdminServiceImpl) AddUser() {
	panic("unimplemented")
}

// BlockUser implements interfaces.AdminService.
func (*AdminServiceImpl) BlockUser() {
	panic("unimplemented")
}

// DeleteUser implements interfaces.AdminService.
func (*AdminServiceImpl) DeleteUser() {
	panic("unimplemented")
}

// FindAllUsers implements interfaces.AdminService.
func (*AdminServiceImpl) FindAllUsers() {
	panic("unimplemented")
}

// FindUser implements interfaces.AdminService.
func (*AdminServiceImpl) FindUser() {
	panic("unimplemented")
}

// Login implements interfaces.AdminService.
func (*AdminServiceImpl) Login() {
	panic("unimplemented")
}

// SearchUser implements interfaces.AdminService.
func (*AdminServiceImpl) SearchUser() {
	panic("unimplemented")
}

// UpdateUser implements interfaces.AdminService.
func (*AdminServiceImpl) UpdateUser() {
	panic("unimplemented")
}

func NewAdminService(repository repository.AdminRepository) service.AdminService {
	return &AdminServiceImpl{}
}
