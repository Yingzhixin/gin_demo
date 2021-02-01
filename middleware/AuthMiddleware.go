package middleware

import (
	"net/http"
	"strings"

	"github.com/Yingzhixin/gin_demo/common"
	"github.com/Yingzhixin/gin_demo/models"
	"github.com/gin-gonic/gin"
)

//AuthMiddleware 用户认证中间件
func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		//验证token格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		//验证token合法性
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		//验证通过之后获取 claim 中的 userID
		userID := claims.UserID
		DB := common.GetDB()
		var user models.User
		DB.First(&user, userID)

		//验证用户是否存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		//用户存在 将 user 的信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
