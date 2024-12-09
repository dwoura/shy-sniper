package svc

import (
	"common"
	"market-monitor/api/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := common.InitGorm(c.DataSource)
	redis, _ := redis.NewRedis(c.Redis)

	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  redis,
	}
}
