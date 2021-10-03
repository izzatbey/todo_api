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

	user := domain.User{
		Fullname: request.Fullname,
		Email:    request.Email,
		Password: request.Password,
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request user.UserUpdateRequest) user.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)
}

func (service UserServiceImpl) FindById(ctx context.Context, userId int) user.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)

}

func (service UserServiceImpl) FindAll(ctx context.Context) []user.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)
}
