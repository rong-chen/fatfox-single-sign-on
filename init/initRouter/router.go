package initRouter

import (
	"fatfox-single-sign-on/core/email"
	"fatfox-single-sign-on/core/token"
	"fatfox-single-sign-on/core/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouterInterface interface {
	InitRouter(*gin.RouterGroup)
}

var RouterList = []RouterInterface{
	new(user.Router),
	new(email.Router),
	new(token.Router),
}

func InitRouter(e *gin.Engine) {
	// 所有其他路由都返回 index.html
	e.StaticFile("/", "./dist/index.html") // 前端网页入口页面
	e.Static("/assets", "./dist/assets")
	r := e.Group("")
	Cors(r)
	for _, routerInterface := range RouterList {
		routerInterface.InitRouter(r)
	}
}
func Cors(router *gin.RouterGroup) {
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // 一天内不再发送预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})
}
