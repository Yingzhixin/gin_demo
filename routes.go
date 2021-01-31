package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yingzhixin/demo/controller"
)

//CollectRoute 处理http
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	return r
}
