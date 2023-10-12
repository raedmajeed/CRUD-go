package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct{
	Engine *gin.Engine
}


func (s *Server) StartServer() {
	s.Engine.Run(":9090")
}

func Start() *Server {
	fmt.Println("engine started")
	router := gin.Default()

	// router.GET("/t", func(ctx *gin.Context) {
	// 	fmt.Println("das")
	// })

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return &Server{
		Engine: router,
	}
}
