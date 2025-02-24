package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	name string
}

func NewUser(name string) *User {
	return &User{
		name: name,
	}
}

func (u *User) GetName() string {
	return u.name
}