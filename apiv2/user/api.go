package user

import (
	"fatfox-single-sign-on/core/user"
	"fatfox-single-sign-on/global"
	"fatfox-single-sign-on/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(c *gin.Context) {
	var lp LoginParams
	err := c.BindJSON(&lp)
	if err != nil {
		c.JSON(200, global.BackResp(400, "参数错误", nil))
		return
	}
	userinfo := user.FindUser("username", lp.Username)
	if ok := utils.ComparePasswords(userinfo.Password, lp.Password); !ok {
		c.JSON(200, global.BackResp(400, "用户名或密码错误", nil))
		return
	}
	//1.验证公司code是否正确 lp.companyId
	//2.根据请求类型获取相对应的临时code lp.ResponseType
	//3.生成临时code有效期1分钟，存到redis中
	code := utils.GenerateRandomCode(32)
	global.RedisDb.Set(c, code, userinfo, 60*time.Second*3)
	c.JSON(200, global.BackResp(302, "", fmt.Sprintf("%s?oAuthCode=%s", lp.Redirect, code)))
}
