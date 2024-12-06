package logic

import (
	"context"
	"doge-arbitrage-system/app/market-monitor/api/internal/svc"
	"doge-arbitrage-system/app/market-monitor/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() (resp *types.Resp, err error) {
	resp = new(types.Resp)
	resp.Msg = "pong"

	return
}
