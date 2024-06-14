package main

import (
	"fatfox-single-sign-on/init/initMigrate"
	"fatfox-single-sign-on/init/initMySql"
	"fatfox-single-sign-on/init/initRedis"
	"fatfox-single-sign-on/init/initRouter"
	"fatfox-single-sign-on/init/initViper"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//初始化配置文件
	c := initViper.InitViper()
	//	初始化数据库
	initMySql.InitMySQL(c)
	initRedis.InitRedis(c)
	initMigrate.InitMigrate()
	g := gin.New()
	initRouter.InitRouter(g)
	err := g.Run(":999")
	if err != nil {
		log.Fatal("run server error" + err.Error())
	}
}
