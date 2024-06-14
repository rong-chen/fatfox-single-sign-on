package global

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

var (
	MySqlDb *gorm.DB
	RedisDb *redis.Client
)

type Model struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func BackResp(code uint, msg string, data interface{}) gin.H {
	// 返回响应
	return gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}

func BackRespData(code uint, data interface{}, total int, page int, pageSize int) gin.H {
	// 返回分页数据
	return gin.H{
		"code":     code,
		"data":     data,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}
}
