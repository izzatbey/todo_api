package repository

import (
	"context"
	"fmt"
	"project_go/app"
	"project_go/model/domain"
	"testing"
	"time"
)

func TestUserInsert(t *testing.T) {
	db := app.NewDB()
	userRepository := NewUserRepository()
	ctx := context.Background()
	user := domain.User{
		Fullname:  "Izzat2",
		Email:     "izzatbey2@gmail.com",
		Password:  "izzatbey2",
		IdRole:    int(1),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, err := userRepository.Insert(ctx, db, user)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(result)
}
