package logic

import (
	"context"
	"github.com/spruceid/siwe-go"

	"github.com/zeromicro/go-zero/core/logx"
	"user/api/internal/svc"
)

type NonceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNonceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NonceLogic {
	return &NonceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NonceLogic) Nonce() (resp string, err error) {
	nonce := siwe.GenerateNonce()
	l.svcCtx.Redis.Setex(nonce, "true", 1800) // 3分钟过期
	return nonce, nil
}
