package dto

import "ginEssential/model"

type UserDto struct {
	Name string `json:name`
	Telephone string `json:telephone`
}

func ToUserDTo(user model.User) UserDto {
	return UserDto{
		Name: user.Name,
		Telephone: user.Telephone,
	}
}

// DTO（Data Transfer Object）数据传输对象。在严格的java EE应用中，
//中间层的组件会将应用底层的状态信息封装成javaBean集，这些JavaBean也被称为DTO。