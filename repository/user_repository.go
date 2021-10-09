package repository

import (
	"context"
	"database/sql"
	"project_go/model/domain"
)

type UserRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, userEmail string) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
}
