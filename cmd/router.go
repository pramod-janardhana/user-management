package main

import (
	"user-management/api"

	"github.com/gin-gonic/gin"
)

func register(router *gin.Engine) {
	userRouter := router.Group("/user")
	{
		userRouter.POST("", api.CreateUser)
		userRouter.GET("/:id", api.GetUser)
	}
}
