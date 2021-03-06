package main

import (
	"github.com/Yingzhixin/gin_demo/controller"
	"github.com/Yingzhixin/gin_demo/middleware"
	"github.com/gin-gonic/gin"
)

//CollectRoute 处理http
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
