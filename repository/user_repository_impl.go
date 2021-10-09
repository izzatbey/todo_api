package repository

import (
	"context"
	"database/sql"
	"errors"
	"project_go/helper"
	"project_go/model/domain"
	"strconv"
	"time"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	sql := "INSERT INTO tb_user(full_name, email, password, forgot_password, id_role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, sql, user.Fullname, user.Email, user.Password, user.ForgotPassword, user.IdRole, time.Now(), time.Now())
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	sql := "UPDATE tb_user set full_name = ?, email = ?, password = ?, forgot_password = ?, id_role = ?, updated_at = ? WHERE id_user = ?"
	_, err := tx.ExecContext(ctx, sql, user.Fullname, user.Email, user.Password, user.ForgotPassword, user.IdRole, user.UpdatedAt, user.Id)
	helper.PanicIfError(err)
	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	sql := "DELETE FROM tb_user WHERE id_user = ?"
	_, err := tx.ExecContext(ctx, sql, user.Id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	sql := "SELECT * FROM tb_user WHERE id_user = ?"
	rows, err := tx.QueryContext(ctx, sql, userId)
	user := domain.User{}
	helper.PanicIfError(err)
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&user.Id, &user.Fullname, &user.Email, &user.Password, &user.ForgotPassword, &user.IdRole, &user.UpdatedAt, &user.UpdatedAt)
		return user, nil
	} else {
		return user, errors.New("ID" + strconv.Itoa(int(userId)) + "Not Found")
	}
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, userEmail string) (domain.User, error) {
	sql := "SELECT * FROM tb_user WHERE email = ?"
	rows, err := tx.QueryContext(ctx, sql, userEmail)
	user := domain.User{}
	helper.PanicIfError(err)
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&user.Id, &user.Fullname, &user.Email, &user.Password, &user.ForgotPassword, &user.IdRole, &user.UpdatedAt, &user.UpdatedAt)
		return user, nil
	} else {
		return user, errors.New("Email" + userEmail + "Not Found")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	sql := "SELECT * FROM tb_user"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()
	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		rows.Scan(&user.Id, &user.Fullname, &user.Email, &user.Password, &user.ForgotPassword, &user.IdRole, &user.UpdatedAt, &user.UpdatedAt)
		users = append(users, user)
	}
	return users
}
