package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Password    string  `json:"-"`
	RawPassword string  `json:"-" gorm:"-"`
	Avatar      string  `json:"avatar"`
	Friends     []User  `json:"fiends,omitempty" gorm:"many2many:user_friends;"`
	Events      []Event `json:"events,omitempty"  gorm:"many2many:user_events;"`
}

func (u *User) makeHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

func (u *User) IsValidPassword(password string) bool {
	hash, err := u.makeHash(password)
	if err != nil {
		log.Fatalln(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.RawPassword != "" {
		password, err := u.makeHash(u.Password)
		if err != nil {
			return err
		}

		u.Password = password
	}

	return nil
}
