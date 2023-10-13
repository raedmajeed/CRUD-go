package controllers

import (
	"raedmajeed/handlers"
	"raedmajeed/util"
)

type UserRouters struct{
	router *ServerStruct
	user *handlers.UserHandler
	jwt *util.JwtUtil
}
 
func (as *UserRouters) URoutes() {
	as.router.r.POST("api/user/register", as.user.RegisterUser)
	as.router.r.POST("api/user/login", as.user.Login)
	as.router.r.GET("api/user/home", as.jwt.ValidateToken("user"),as.user.Home)
}

func NewUserRoute(a *handlers.UserHandler, server *ServerStruct, jwt *util.JwtUtil) *UserRouters {
	return &UserRouters{
		router: server,
		user: a,
		jwt: jwt,
	}
}
