package router

import (
	"microservice1/handler/user"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	userGroup := r.Group("user")

	userGroup.POST("/verify", user.VerifyUser)
	userGroup.POST("/save", user.SaveUser)
}