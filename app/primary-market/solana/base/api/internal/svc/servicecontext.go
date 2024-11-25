package svc

import (
	"base/api/internal/config"
	"base/rpc/baseservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	BaseService baseservice.BaseService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		BaseService: baseservice.NewBaseService(zrpc.MustNewClient(c.Base)),
	}
}
