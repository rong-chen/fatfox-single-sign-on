package token

import "github.com/gin-gonic/gin"

type Router struct{}

func (Router) InitRouter(r *gin.RouterGroup) {
	router := r.Group("/token")
	{
		router.GET("/refresh", UpdateToken)
	}
}
