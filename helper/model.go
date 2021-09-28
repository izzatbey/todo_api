package helper

import (
	"project_go/model/domain"
	"project_go/model/web/user"
)

func ToUserResponse(user domain.User) user.UserResponse {
	return user.UserResponse{}
}
