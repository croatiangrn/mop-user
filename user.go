package mop_user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID        int    `gorm:"primaryKey;" json:"id"`
	FirstName string `gorm:"not null;type:varchar(70);" json:"first_name"`
	LastName  string `gorm:"not null;type:varchar(70);" json:"last_name"`
	Email     string `gorm:"not null;uniqueIndex:ux_email;type:varchar(255);" json:"email"`
	Password  string `gorm:"not null;type:varchar(255)" json:"password"`
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

	hashedPass, hashedPassErr := hashPassword(data.Password)
	if hashedPassErr != nil {
		return hashedPassErr
	}

	if err := u.db.Debug().Exec(query, data.FirstName, data.LastName, data.Email, hashedPass).Error; err != nil {
		log.Printf("error while registering user: %v\n", err)
		if isUniqueConstraintViolation(err) {
			return ErrEmailTaken
		}
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

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordSalt+password))
	return err == nil
}
