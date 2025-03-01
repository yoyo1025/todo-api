package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	name string
	email string
}

func NewUser(name, email string) *User {
	return &User{
		name: name,
		email: email,
	}
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}