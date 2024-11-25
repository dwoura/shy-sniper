package svc

import (
	"base/rpc/internal/config"
	"github.com/gagliardetto/solana-go/rpc"
)

type ServiceContext struct {
	Config       config.Config
	SolanaClient *rpc.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		SolanaClient: rpc.New(rpc.MainNetBeta_RPC),
	}
}
