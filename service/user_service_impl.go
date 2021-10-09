package service

import (
	"context"
	"database/sql"
	"project_go/helper"
	"project_go/model/domain"
	"project_go/model/web/user"
	"project_go/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepositoryImpl
	DB             *sql.DB
}

func (service *UserServiceImpl) Create(ctx context.Context, request user.UserCreateRequest) user.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)

	users := domain.User{
		Fullname:       request.Fullname,
		Email:          request.Email,
		Password:       request.Password,
		ForgotPassword: request.ForgotPassword,
		IdRole:         request.RoleId,
	}

	userResult := service.UserRepository.Insert(ctx, tx, users)

	return helper.ToUserResponse(userResult)
}

func (service *UserServiceImpl) Update(ctx context.Context, request user.UserUpdateRequest) user.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)

	users, err := service.UserRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	users.Fullname = request.Fullname
	users.Email = request.Email
	users.Password = request.Password
	users.ForgotPassword = request.ForgotPassword
	users.IdRole = request.RoleId

	userResult := service.UserRepository.Update(ctx, tx, users)

	return helper.ToUserResponse(userResult)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)

	users, err := service.UserRepository.FindById(ctx, tx, userId)
	helper.PanicIfError(err)

	service.UserRepository.Delete(ctx, tx, users)
}

func (service UserServiceImpl) FindById(ctx context.Context, userId int) user.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)

	users, err := service.UserRepository.FindById(ctx, tx, userId)
	helper.PanicIfError(err)

	return helper.ToUserResponse(users)
}

func (service UserServiceImpl) FindAll(ctx context.Context) []user.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)

	return helper.ToUserResponses(users)
}
