package model

import (
	"gofly/utils"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `json:"id" gorm:"primary key"`
	CreatedAt Time   `json:"create_at"`
	UpdatedAt Time   `json:"update_at"`
	Name      string `json:"name" gorm:"size:64;not null;comment:昵称"`
	RealName  string `json:"real_name" gorm:"size:128;comment:真实姓名"`
	Avatar    string `json:"avatar" gorm:"size:255;comment:头像"`
	Mobile    string `json:"mobile" gorm:"size:11;comment:手机号"`
	Email     string `json:"email" gorm:"128;comment:邮箱"`
	Password  string `json:"-" gorm:"128;not null;comment:密码"`
}

func (u *User) encrypt() error {
	stHash, err := utils.Encrypt(u.Password)
	if err == nil {
		u.Password = stHash
	}
	return err
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return u.encrypt()
}

// LoginUser 用户登录信息 // ========================================
type LoginUser struct {
	ID   uint
	Name string
}
