package models

import (
	"time"

	"vinpel/golang-sample-api-jwt/common/dto"

	"golang.org/x/crypto/bcrypt"
)

// User datan, not exposed
type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name"`
	Username  string     `json:"username" gorm:"unique"`
	Email     string     `json:"email" gorm:"unique"`
	Password  string     `json:"password"`
	Empreinte string     `json:"empreinte"`
}

func (user *User) CreateFromQuery(query *dto.UserRequest) {
	user.Name = query.Name
	user.Username = query.Username
	user.Email = query.Email
	user.Password = query.Password
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
