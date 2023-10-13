package di

import (
	"raedmajeed/controllers"
	"raedmajeed/db"
	"raedmajeed/handlers"
	"raedmajeed/repository"
	"raedmajeed/service"
	"raedmajeed/util"
)

func Init() *controllers.ServerStruct{
	db := db.ConnectDatabase()
	jwt := util.NewJwtUtil()
	adminRepository := repository.NewAdminRepository(db)
	userRepository := repository.NewUserRepository(db)
	
	adminService := service.NewAdminService(adminRepository, jwt)
	userService := service.NewUserService(userRepository, jwt)
	
	server := controllers.NewHTTPServer()
	
	adminHandlers := handlers.NewAdminHandler(adminService)
	userHandlers := handlers.NewUserHandler(userService)
	
	userRoutes := controllers.NewUserRoute(userHandlers, server, jwt)
	adminRoutes := controllers.NewAdminRoute(adminHandlers, server, jwt)
	adminRoutes.Routes()
	userRoutes.URoutes()
	return server
}