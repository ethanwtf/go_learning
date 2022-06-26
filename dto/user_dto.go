package dto

import "gin_vue_project/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"Telephone"`
}

// 返回用户数据不能带密码
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Phone,
	}
}
