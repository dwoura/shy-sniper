package logic

import (
	"base/rpc/baseservice"
	"context"

	"base/api/internal/svc"
	"base/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBalanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBalanceLogic {
	return &GetBalanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBalanceLogic) GetBalance(req *types.GetBalanceReq) (resp *types.GetBalanceResp, err error) {
	out, err := l.svcCtx.BaseService.GetBalance(l.ctx, &baseservice.GetBalanceReq{
		Address: req.Address,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetBalanceResp{
		Balance: out.Balance,
	}, nil
}
