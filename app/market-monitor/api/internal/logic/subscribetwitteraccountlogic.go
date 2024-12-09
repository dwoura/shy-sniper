package logic

import (
	"context"

	"market-monitor/api/internal/svc"
	"market-monitor/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubscribeTwitterAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubscribeTwitterAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubscribeTwitterAccountLogic {
	return &SubscribeTwitterAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubscribeTwitterAccountLogic) SubscribeTwitterAccount(req *types.SubscribeReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line

	return
}
