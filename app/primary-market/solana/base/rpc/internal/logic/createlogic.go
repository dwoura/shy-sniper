package logic

import (
	"context"

	"base/rpc/base"
	"base/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *base.CreateReq) (*base.CreateResp, error) {
	// todo: add your logic here and delete this line

	return &base.CreateResp{}, nil
}
