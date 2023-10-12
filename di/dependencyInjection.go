package di

import (
	"fmt"
	"raedmajeed/controllers"
	"raedmajeed/db"
	"raedmajeed/repository"
	"raedmajeed/service"
)

func Init() {
	db := db.ConnectDatabase()
	// userRepository := repository.NewUserRepository(db)
	adminRepository := repository.NewAdminRepository(db)
	// userService := service.NewAdminServiceImpl(userRepository)
	adminService := service.NewAdminService(adminRepository)
	// fmt.Println(adminService)
	server := controllers.Start()
	adminHandler := controllers.NewAdminHandler(adminService, server)
	adminHandler.Handler()

	userHandler := controllers.NewUserHandler(adminService, server)
	userHandler.UHandler()
	fmt.Println(adminHandler)
	
	server.StartServer()
}