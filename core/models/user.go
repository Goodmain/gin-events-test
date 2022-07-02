package models

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Name               string  `json:"name"`
	Email              string  `json:"email"`
	Password           string  `json:"-"`
	RawPassword        string  `json:"-" gorm:"-"`
	ResetPasswordToken string  `json:"-"`
	Avatar             string  `json:"avatar"`
	Friends            []User  `json:"fiends,omitempty" gorm:"many2many:user_friends;"`
	Events             []Event `json:"events,omitempty" gorm:"many2many:user_events;"`
	//FriendedBy         []User  `json:"friended_by,omitempty" gorm:"many2many:user_friends;ForeignKey:id;References:friend_id"`
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

func (u *User) GenerateResetPasswordToken() {
	h := md5.New()
	h.Write([]byte(strings.ToLower(u.Email + u.Password)))
	u.ResetPasswordToken = hex.EncodeToString(h.Sum(nil))
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
