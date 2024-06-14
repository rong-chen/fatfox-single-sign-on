package email

import "github.com/gin-gonic/gin"

type Router struct {
}

func (Router) InitRouter(group *gin.RouterGroup) {
	r := group.Group("email")
	{
		r.POST("/send", Send)
	}
}
