package email

import (
	"context"
	"fatfox-single-sign-on/core/user"
	"fatfox-single-sign-on/global"
	"fatfox-single-sign-on/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func Send(c *gin.Context) {
	var p Params
	err := c.BindJSON(&p)
	if err != nil {
		c.JSON(200, global.BackResp(400, "参数错误", err))
		return
	}
	ctx := context.Background()
	res, _ := global.RedisDb.Get(ctx, p.Email).Result()
	if res != "" {
		c.JSON(200, global.BackResp(400, "请勿重复发起验证码", err))
		return
	}
	info := user.FindUser("email", p.Email)
	if info.Id != uuid.Nil {
		c.JSON(200, global.BackResp(400, "邮箱已注册", err))
		return
	}
	var m Mailer
	m.T = p.Email
	m.F = "1416307833@qq.com"
	m.C = "1416307833@qq.com"
	m.Account = "1416307833@qq.com"
	m.Password = "yetlmmkncinzhdjj"
	code := utils.GenerateRandomCode(6)
	m.HtmlBody = fmt.Sprintf(HtmlBody, code)
	err = global.RedisDb.Set(ctx, p.Email, code, 60*1*time.Second).Err()
	if err != nil {
		c.JSON(200, global.BackResp(400, "网络异常", err))
		return
	}
	// 发送邮件
	go func(email string, m Mailer) {
		err = m.SendEmail()
		if err != nil {
			c.JSON(200, global.BackResp(400, "发送验证码失败", err))
			return
		}
	}(p.Email, m)
	c.JSON(200, global.BackResp(200, "发送成功", nil))
}
