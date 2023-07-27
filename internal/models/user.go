package models

import (
	"github.com/pkg/errors"
	"github.com/shinhagunn/todo-backend/pkg/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Model

	UID      string `gorm:"type:character varying(13);not null"`
	Email    string `gorm:"type:character varying;not null;unique"`
	Password string `gorm:"type:character varying;not null"`
}

func (u User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.UID = util.GenerateUID()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "hash password fail")
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return errors.Wrap(err, "compare hash and password fail")
}
