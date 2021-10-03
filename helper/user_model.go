package helper

import (
	"project_go/model/domain"
	"project_go/model/web/user"
)

func ToUserResponse(value domain.User) user.UserResponse {
	return user.UserResponse{
		Id:       value.Id,
		Fullname: value.Fullname,
		Email:    value.Email,
	}
}
