package controller

import (
	"log"
	"net/http"

	"github.com/Yingzhixin/gin_demo/common"
	"github.com/Yingzhixin/gin_demo/dto"
	"github.com/Yingzhixin/gin_demo/models"
	"github.com/Yingzhixin/gin_demo/response"
	"github.com/Yingzhixin/gin_demo/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//Register 注册用户
func Register(ctx *gin.Context) {

	DB := common.GetDB()
	//获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	//如果名称没有传,那么给一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	//判断手机号是否已经被注册
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号已经被占用")
		return
	}

	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "密码加密错误")
		return
	}
	newUser := models.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)

	//返回结果
	response.Success(ctx, nil, "注册成功")
}

//Login 用户登陆
func Login(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	//判断手机号是否存在
	var user models.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(ctx, nil, "密码错误")
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "系统错误,用户token生成异常.")
		log.Printf("token generate error: %v", err)
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token}, "登陆成功")
}

//Info 用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(models.User))}, "获取用户信息成功")
}

//isTelephoneExist 检测手机号码是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user models.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
