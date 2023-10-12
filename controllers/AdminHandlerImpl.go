package controllers

import (
	"raedmajeed/controllers/interfaces"
	service "raedmajeed/service/interfaces"

	"github.com/gin-gonic/gin"
)

type AdminHandlerImpl struct{
	router *gin.Engine
}

func (ah *AdminHandlerImpl) Handler() {
	ah.router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H {
			"message": "hi",
		})
	})
	// ah.router.GET("/test", func(ctx *gin.Context) {
	// 	fmt.Println("test")
	// })
}

func NewAdminHandler(si service.AdminService, r *Server) interfaces.AdminHandlers {
	return &AdminHandlerImpl{
		router: r.Engine,
	}
}
