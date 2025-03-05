package dto

import "todo-api/domain/model"

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

type UserResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func ToUserResponse(u *model.User) UserResponse {
	return UserResponse {
		ID: uint(u.GetID()),
		Name: u.GetName(),
		Email: u.GetEmail(),
	}
}