package api

import (
	"log"
	"user-management/domain"
	"user-management/internal/controller"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	request := domain.User{}
	c.Bind(&request)
	// TODO: add request validation
	userController := controller.NewUserController()
	err := userController.Create(&request)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"err":  "",
		"data": "user created",
	})
}

func GetUser(c *gin.Context) {
	userId := c.Params.ByName("id")
	userController := controller.NewUserController()
	toReaturn, err := userController.Get(userId)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"err":  "",
		"data": toReaturn,
	})
}
