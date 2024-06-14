package user

import "github.com/gin-gonic/gin"

type Router struct {
}

func (Router) InitRouter(group *gin.RouterGroup) {
	r := group.Group("user")
	{
		r.POST("/register", Register)
		r.POST("/login", Login)
	}
}
