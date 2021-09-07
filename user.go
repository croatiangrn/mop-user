package mop_user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID        int    `gorm:"primaryKey;" json:"id"`
	FirstName string `gorm:"not null;" json:"first_name"`
	LastName  string `gorm:"not null;" json:"last_name"`
	Email     string `gorm:"uniqueIndex:ux_email;" json:"email"`
	Password  string `gorm:"not null;" json:"password"`
	db        *gorm.DB
}

func (u *User) TableName() string {
	return "users"
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (u *User) Register(data UserRegistration) error {
	query := `INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)`

	if err := u.db.Debug().Exec(query, data.FirstName, data.LastName, data.Email, data.Password).Error; err != nil {
		log.Printf("error while registering user: %v\n", err)
		return ErrInternal
	}

	userID, err := getLastInsertedID(u.db)
	if err != nil {
		return err
	}

	u.ID = userID
	u.FirstName = data.FirstName
	u.LastName = data.LastName
	u.Email = data.Email
	hashedPass, hashedPassErr := hashPassword(data.Password)
	if hashedPassErr != nil {
		return hashedPassErr
	}

	u.Password = hashedPass

	return nil
}

func hashPassword(password string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(passwordSalt+password), 5)
	if err != nil {
		log.Printf("error while hashing the password: %v\n", err)
		return string(hashBytes), ErrHashingPassword
	}

	return string(hashBytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordSalt+password))
	return err == nil
}
