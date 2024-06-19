package token

import (
	"fatfox-single-sign-on/global"
	"fatfox-single-sign-on/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func Refresh(c *gin.Context) {
	rToken := c.Query("rToken")
	if rToken == "" {
		c.JSON(200, global.BackResp(400, "参数不能为空", nil))
		return
	}

	p, err := utils.ParseJWT(rToken)
	if err != nil {
		c.JSON(200, global.BackResp(400, "无效参数", nil))
		return
	}

	//判断token类型是否是refresh token
	if p.Type != utils.RefreshToken {
		c.JSON(200, global.BackResp(400, "token类型错误", nil))
		return
	}

	// 验证sessionId是否有效
	sessionId := global.RedisDb.Get(c, p.Params.Id.String()).Val()
	if sessionId != p.SessionId {
		c.JSON(200, global.BackResp(400, "用户登录信息已失效", nil))
		return
	}
	// 生成新的token
	aToken, err := utils.GenerateJWT(p.Params, utils.AccessToken, "", time.Now().Add(15*time.Minute))
	if err != nil {
		c.JSON(200, global.BackResp(400, "生成失败", nil))
		return
	}
	c.JSON(200, global.BackResp(200, "", aToken))
}

func validToken(c *gin.Context) {
	t, err := utils.ParseJWT(c.Query("aToken"))
	if (err != nil) || (t.Params.Id == uuid.Nil) {
		c.JSON(200, global.BackResp(400, "无效参数", nil))
		return
	}
	//验证token类型
	if t.Type != utils.RefreshToken {
		c.JSON(200, global.BackResp(400, "无效参数", nil))
		return
	}
	c.JSON(200, global.BackResp(200, "", nil))
}

func ValidTick(c *gin.Context) {
	// 验证ticket
	t, err := utils.ParseJWT(c.Query("ticket"))
	if (err != nil) || (t.Params.Id == uuid.Nil) {
		c.JSON(200, global.BackResp(400, "无效参数", nil))
		return
	}

	//验证ticket类型
	if t.Type != utils.TicketToken {
		c.JSON(200, global.BackResp(400, "无效参数", nil))
		return
	}

	sessionId := utils.GenerateRandomCode(8)
	rToken, err := utils.GenerateJWT(t.Params, utils.RefreshToken, sessionId, time.Now().Add(time.Hour*24*7))
	aToken, err := utils.GenerateJWT(t.Params, utils.AccessToken, "", time.Now().Add(15*time.Minute))
	if err != nil {
		c.JSON(200, global.BackResp(500, "生成rToken失败", nil))
		return
	}

	var data = make(map[string]interface{})
	data["rToken"] = rToken
	data["aToken"] = aToken
	data["userinfo"] = t.Params
	c.JSON(200, global.BackResp(200, "", data))
}
