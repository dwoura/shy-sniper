package logic

import (
	"context"

	"market-monitor/api/internal/svc"
	"market-monitor/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubscribeBinanceAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubscribeBinanceAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubscribeBinanceAccountLogic {
	return &SubscribeBinanceAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubscribeBinanceAccountLogic) SubscribeBinanceAccount(req *types.SubscribeReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line

	return
}
