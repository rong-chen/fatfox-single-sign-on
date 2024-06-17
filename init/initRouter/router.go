package initRouter

import (
	"fatfox-single-sign-on/core/email"
	"fatfox-single-sign-on/core/token"
	"fatfox-single-sign-on/core/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
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

	// Fallback route to serve index.html for Vue Router history mode
	e.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	r := e.Group("")
	r.Any("/api/*any", reverseProxy("127.0.0.1"))

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

// 反向代理的处理器
func reverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析目标地址
		targetUrl, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid target URL"})
			return
		}

		// 创建反向代理
		proxy := httputil.NewSingleHostReverseProxy(targetUrl)
		proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
			http.Error(rw, "Proxy error: "+err.Error(), http.StatusBadGateway)
		}

		// 重写请求 URL
		c.Request.URL.Host = targetUrl.Host
		c.Request.URL.Scheme = targetUrl.Scheme
		c.Request.Header.Set("X-Forwarded-Host", c.Request.Host)
		c.Request.Host = targetUrl.Host

		// 执行反向代理
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
