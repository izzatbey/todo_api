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

func ToUserResponses(users []domain.User) []user.UserResponse {
	var userResponses []user.UserResponse
	for _, useridx := range users {
		userResponses = append(userResponses, ToUserResponse(useridx))
	}
	return userResponses
}
