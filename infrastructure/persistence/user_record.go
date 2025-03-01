package persistence

import "gorm.io/gorm"

type UserRecord struct {
	gorm.Model
	Name string `gorm:"columun:name"`
	Email string `gorm:"columun:email"`
}