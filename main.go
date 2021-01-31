package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yingzhixin/demo/common"
)

func main() {
	common.InitDB()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}
