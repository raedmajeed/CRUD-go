package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ServerStruct struct {
	r *gin.Engine
}

func (s *ServerStruct) StartServer() {
	s.r.Run(":9090")
}

func NewHTTPServer() *ServerStruct{
	router := gin.Default()

	router.GET("/tes", func(ctx *gin.Context) {
		fmt.Println("test")
	})

	return &ServerStruct {
		r: router,
	}
}