package company

import (
	"fatfox-single-sign-on/global"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RegisterAuthority(c *gin.Context) {
	var params RegisterAuthorityParams
	err := c.BindJSON(&params)
	if err != nil {
		c.JSON(200, global.BackResp(500, "参数错误", nil))
		return
	}
	//查询redis确定是公司id申请的code，接着拿到公司id去校验参数secret是否正确
	result, err := global.RedisDb.Get(c, params.TemporaryCode).Result()
	if err != nil {
		c.JSON(200, global.BackResp(500, "参数错误", nil))
		return
	}
	fmt.Println(result)
	//校验过程没写，代补
	//校验成功后，通过这个用户Id和公司id去创建一个account_token
}
