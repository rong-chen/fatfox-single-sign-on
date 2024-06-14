package user

import (
	"fatfox-single-sign-on/global"
	"fmt"
)

func FindUser(key string, val string) (info *Info) {
	global.MySqlDb.Where(fmt.Sprintf("%s = ?", key), val).Find(&info)
	return
}

func Create(info *Info) error {
	return global.MySqlDb.Create(&info).Error
}
