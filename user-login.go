package mop_user

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

type UserLoginRequestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	db       *gorm.DB
}

func NewUserLoginRequestData(db *gorm.DB) *UserLoginRequestData {
	return &UserLoginRequestData{db: db}
}

func (l *UserLoginRequestData) ProcessLogin() (*UserLoginResponseData, error) {
	if len(l.Email) == 0 || len(l.Password) < PassMinLength {
		return nil, ErrInvalidMailOrPassword
	}

	user := User{}
	query := `SELECT * FROM users WHERE email = ?`

	if err := l.db.Debug().Raw(query, l.Email).Scan(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidMailOrPassword
		}

		log.Printf("error occurred while getting user from db: %v\n", err)
		return nil, ErrInternal
	}

	if !CheckPasswordHash(l.Password, user.Password) {
		return nil, ErrInvalidMailOrPassword
	}

	response := UserLoginResponseData{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	return &response, nil
}

type UserLoginResponseData struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Token     string `json:"token"`
}
