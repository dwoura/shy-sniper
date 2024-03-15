package logic

import (
	"context"

	"shy-sniper/shysniperbot/internal/svc"
	"shy-sniper/shysniperbot/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShysniperbotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShysniperbotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShysniperbotLogic {
	return &ShysniperbotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShysniperbotLogic) Shysniperbot(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
