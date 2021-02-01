package main

import (
	"github.com/Yingzhixin/gin_demo/common"
	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}
