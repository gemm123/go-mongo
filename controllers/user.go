package controllers

import (
	"net/http"

	"github.com/gemm123/go-mongo/models"
	"github.com/gemm123/go-mongo/services"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *userController {
	return &userController{userService}
}

func (uc *userController) PostUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.userService.CreateUser(user); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

func (uc *userController) GetAllUser(c *gin.Context) {
	users, err := uc.userService.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (uc *userController) GetUserByName(c *gin.Context) {
	name := c.Param("name")
	user, err := uc.userService.GetUserByName(name)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (uc *userController) UpdateUser(c *gin.Context) {
	name := c.Param("name")

	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.userService.UpdateUser(name, newUser); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

func (uc *userController) DeleteUser(c *gin.Context) {
	name := c.Param("name")

	if err := uc.userService.DeleteUser(name); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
