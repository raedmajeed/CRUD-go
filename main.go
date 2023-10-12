package main

import (
	"raedmajeed/di"

	"github.com/gin-gonic/gin"
)
var Router *gin.Engine


func main() {

	di.Init()

}