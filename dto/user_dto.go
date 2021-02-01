package dto

import "github.com/Yingzhixin/gin_demo/models"

//UserDto 用户数据传输对象
type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

//ToUserDto 数据对象转化
func ToUserDto(user models.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
