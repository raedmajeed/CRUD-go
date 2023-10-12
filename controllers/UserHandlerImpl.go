package controllers

import (
	"raedmajeed/controllers/interfaces"
	service "raedmajeed/service/interfaces"

	"github.com/gin-gonic/gin"
)

type UserHandlerImpl struct{
	router *gin.Engine
}

func (ah *UserHandlerImpl) UHandler() {
	ah.router.GET("/test2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H {
			"message": "hisdas",
		})
	})
}

func NewUserHandler(si service.AdminService, r *Server) interfaces.UserHandlers {
	return &UserHandlerImpl{
		router: r.Engine,
	}
}
