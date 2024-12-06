package svc

import (
	"doge-arbitrage-system/app/market-monitor/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
	}
}
