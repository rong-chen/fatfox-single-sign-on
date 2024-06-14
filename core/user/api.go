package user

import (
	"context"
	"fatfox-single-sign-on/global"
	"fatfox-single-sign-on/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func Login(c *gin.Context) {

	type Params struct {
		Username string `json:"username" form:"username"  binding:"required"`
		Password string `json:"password" form:"password"  binding:"required"`
		Redirect string `json:"redirect" form:"redirect"  binding:"required"`
	}
	var p Params
	err := c.BindJSON(&p)
	if err != nil {
		c.JSON(200, global.BackResp(400, "参数不正确", err.Error()))
		return
	}
	user := FindUser("username", p.Username)
	if user.Id == uuid.Nil {
		c.JSON(200, global.BackResp(400, "账号密码错误", nil))
		return
	}
	if ok := utils.ComparePasswords(user.Password, p.Password); !ok {
		c.JSON(200, global.BackResp(400, "账号密码错误", nil))
		return
	}
	params := utils.Params{
		Id:       user.Id,
		Email:    user.Email,
		Nickname: user.Nickname,
		Username: user.Username,
		Phone:    user.Phone,
		Avatar:   user.Avatar,
	}
	aToken, _ := utils.GenerateJWT(params, utils.AccessToken, time.Now().Add(15*time.Minute))
	rToken, _ := utils.GenerateJWT(params, utils.RefreshToken, time.Now().Add(24*time.Hour*30))
	resp := make(map[string]any)
	resp["redirect"] = fmt.Sprintf("%s?atoken=%s&rtoken=%s", p.Redirect, aToken, rToken)
	c.JSON(200, global.BackResp(302, "", resp))
}

func Register(c *gin.Context) {
	// 注册用户
	var cp createParams
	err := c.BindJSON(&cp)
	if err != nil {
		c.JSON(200, global.BackResp(400, err.Error(), nil))
		return
	}
	ctx := context.Background()
	code, _ := global.RedisDb.Get(ctx, cp.Email).Result()
	global.RedisDb.Del(ctx, cp.Email)
	if code != cp.Code {
		c.JSON(200, global.BackResp(400, "验证码错误", nil))
		return
	}

	password, _ := utils.HashPassword(cp.Password)
	var info = &Info{
		Email:    cp.Email,
		Password: password,
		Phone:    cp.Phone,
		Username: cp.Username,
		Nickname: cp.Nickname,
	}

	info.Id, _ = uuid.NewUUID()
	err = Create(info)
	if err != nil {
		c.JSON(200, global.BackResp(400, err.Error(), nil))
		return
	}
	c.JSON(200, global.BackResp(200, "注册成功", nil))
}
