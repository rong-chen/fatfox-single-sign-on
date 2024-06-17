package user

import (
	"fatfox-single-sign-on/global"
	"time"
)

type Info struct {
	Username     string     `json:"username" form:"username" gorm:"username" `
	Password     string     `json:"password"  form:"password" gorm:"password"`
	Phone        string     `json:"phone"  form:"phone" gorm:"phone"`
	Email        string     `json:"email;not null;unique"  form:"email" gorm:"email"`
	Nickname     string     `json:"nickname"  form:"nickname" gorm:"nickname"`
	Avatar       string     `json:"avatar"  form:"avatar" gorm:"avatar"`
	Disabled     bool       `json:"disabled" form:"disabled" gorm:"disabled"`
	DisabledTime *time.Time `json:"disabledTime" form:"disabledTime" gorm:"disabledTime"`
	global.Model
}

type createParams struct {
	Username string `json:"username" form:"username" gorm:"username" binding:"required"`
	Password string `json:"password"  form:"password" gorm:"password" binding:"required"`
	Phone    string `json:"phone"  form:"phone" gorm:"phone" binding:"required"`
	Email    string `json:"email"  form:"email" gorm:"email;not null;unique" binding:"required"`
	Nickname string `json:"nickname"  form:"nickname" gorm:"nickname" binding:"required"`
	Code     string `json:"code"  form:"code" gorm:"code" binding:"required"`
}

func (Info) TableName() string {
	return "user_info"
}
