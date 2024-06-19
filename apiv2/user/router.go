package user

import "github.com/gin-gonic/gin"

type Router struct {
}

func (receiver Router) InitRouter(router *gin.RouterGroup) {
	r := router.Group("v2").Group("user")
	{
		// 登录
		r.POST("/login", Login)
	}
}
