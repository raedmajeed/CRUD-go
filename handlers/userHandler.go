package handlers

import (
	"net/http"
	"raedmajeed/dto"
	"raedmajeed/entity"
	"raedmajeed/service/interfaces"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	ctx *gin.Context
	user interfaces.UserService
}

func (ah *UserHandler) RegisterUser(c *gin.Context) {
	user := &entity.User{}
	c.BindJSON(user)
	err, user := ah.user.RegisterUser(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message" : err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H {
		"data": user,
	})
}


func (ah *UserHandler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"message": "You are logged in",
	})
}

func (ah *UserHandler) Login(c *gin.Context) {
	LoginRequest := &dto.LoginRequest{}
	c.BindJSON(LoginRequest)

	err, token := ah.user.Login(LoginRequest)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
				"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusAccepted, gin.H {
		"token": token,
	})
}

func NewUserHandler(userService interfaces.UserService) *UserHandler {
	return &UserHandler{
		user: userService,
	}
}