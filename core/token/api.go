package token

import (
	"fatfox-single-sign-on/core/user"
	"fatfox-single-sign-on/global"
	"fatfox-single-sign-on/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func RegisterSecretKey(c *gin.Context) {
	//
}

func UpdateToken(c *gin.Context) {
	// 刷新token逻辑
	rToken := c.GetHeader("Refresh-Token")
	if rToken == "" {
		c.JSON(200, global.BackResp(400, "Refresh-Token不能为空", nil))
		return
	}
	token, err := utils.ParseJWT(rToken)
	if err != nil {
		c.JSON(200, global.BackResp(400, "Refresh-Token有误", nil))
		return
	}
	if token.Type != utils.RefreshToken {
		c.JSON(200, global.BackResp(400, "Refresh-Token有误", nil))
		return
	}
	userId := token.Params.Id.String()
	userinfo := user.FindUser("id", userId)
	if userinfo.Id == uuid.Nil {
		c.JSON(200, global.BackResp(400, "Refresh-Token有误", nil))
		return
	}

	params := utils.Params{
		Id:       userinfo.Id,
		Email:    userinfo.Email,
		Nickname: userinfo.Nickname,
		Username: userinfo.Username,
		Phone:    userinfo.Phone,
		Avatar:   userinfo.Avatar,
	}

	newToken, _ := utils.GenerateJWT(params, utils.AccessToken, time.Now().Add(time.Minute*15))
	resp := make(map[string]string)
	resp["access_token"] = newToken
	c.JSON(200, global.BackResp(200, "", resp))
}
