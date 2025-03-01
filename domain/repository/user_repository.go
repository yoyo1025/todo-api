package repository

import "todo-api/domain/model"

type IUserRepository interface {
	FindByEmail(email string) (*model.User, error)
	SignUp(name, email string) (*model.User, error)
}