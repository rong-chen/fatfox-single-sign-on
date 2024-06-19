package token

import "github.com/gin-gonic/gin"

type Router struct {
}

func (Router) InitRouter(router *gin.RouterGroup) {
	r := router.Group("token")
	{
		r.GET("/validTick", ValidTick)
		r.GET("/refresh", Refresh)
		r.GET("/validToken", validToken)
	}
}
