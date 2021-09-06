package mop_user

import "gorm.io/gorm"

type User struct {
	ID        int    `gorm:"primaryKey;" json:"id"`
	FirstName string `gorm:"not null;" json:"first_name"`
	LastName  string `gorm:"not null;" json:"last_name"`
	Email     string `gorm:"uniqueIndex:ux_email;" json:"email"`
	Password  string `gorm:"not null;" json:"password"`
	db        *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}
