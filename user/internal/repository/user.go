package repository

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID         uint   `gorm:"primarykey"`
	UserName       string `gorm:"unique"`
	NickName       string
	PasswordDigest string
}

const (
	PassWordCost = 12 // 密码加密难度
)

// 加密密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}
