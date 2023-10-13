package handlers

import (
	"log"
	"net/http"
	"raedmajeed/dto"
	"raedmajeed/entity"
	"raedmajeed/service/interfaces"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	ctx *gin.Context
	admin interfaces.AdminService
}

func (ah *AdminHandler) AddUser(c *gin.Context) {
	user := &entity.User{}
	c.BindJSON(user)
	err, user := ah.admin.AddUser(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message" : err,
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H {
		"data": user,
	})
}

func (ah *AdminHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	user := &entity.User{}
	err = c.BindJSON(user)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
				"message": err,
		})
		return
	}

	err , user = ah.admin.UpdateUser(idInt, *user)
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
					"message": "Unable to update user",
			})
			return
	}
	
	c.JSON(http.StatusCreated, gin.H {
		"data": user,
	})
}

func (ah *AdminHandler) SearchUser(c *gin.Context) {
	str := c.Query("str")
	if str == "" {
		log.Println("Search string is empty")
		c.JSON(http.StatusBadRequest, gin.H {
			"message": "Search string is empty",
		})
		return
	}

	err, users := ah.admin.SearchUser(str)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err,
		})
		return
	}

	log.Println("Search string is empty")
		c.JSON(http.StatusFound, gin.H {
			"data": users,
		})
		return
}

func (ah *AdminHandler) FindUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID",
		})
		return
	}

	err, user := ah.admin.FindUser(userID)
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
					"message": "Unable to Find user",
			})
			return
	}
	c.JSON(http.StatusCreated, gin.H {
		"data": user,
	})
}

func (ah *AdminHandler) FindAllUsers(c *gin.Context) {
	err, users := ah.admin.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message" : "Unable To Find All Users",
		})
		return
	}

	c.JSON(http.StatusFound, gin.H {
		"data": users,
	})
}

func (ah *AdminHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID",
		})
		return
	}
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
				"message": "Invalid user data",
		})
		return
	}

	err , user := ah.admin.DeleteUser(idInt)
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
					"message": "Unable to update user",
			})
			return
	}
	
	c.JSON(http.StatusCreated, gin.H {
		"data": user,
	})
}

func (ah *AdminHandler) BlockUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID",
		})
		return
	}
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
				"message": "Invalid user data",
		})
		return
	}

	err , user := ah.admin.BlockUser(idInt)
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
					"message": "Unable to update user",
			})
			return
	}
	
	c.JSON(http.StatusCreated, gin.H {
		"data": user,
	})
}

func (ah *AdminHandler) UnBlockUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID",
		})
		return
	}
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
				"message": "Invalid user data",
		})
		return
	}

	err , user := ah.admin.UnBlockUser(idInt)
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
					"message": "Unable to update user",
			})
			return
	}
	
	c.JSON(http.StatusCreated, gin.H {
		"data": user,
	})
}

func (ah *AdminHandler) Login(c *gin.Context) {

	LoginRequest := &dto.LoginRequest{}
	c.BindJSON(LoginRequest)

	err, token := ah.admin.Login(LoginRequest)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
				"message": "Invalid Login Details",
		})
		return
	}
	
	c.JSON(http.StatusAccepted, gin.H {
		"token": token,
	})
}

func NewAdminHandler(adminService interfaces.AdminService) *AdminHandler {
	return &AdminHandler{
		admin: adminService,
	}
}