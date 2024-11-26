package logic

import (
	"context"
	"errors"
	"fmt"
	"user/common"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	if req.Username == "admin" && req.Password == "123456" { // todo: 修改判断
		auth := l.svcCtx.Config.Auth
		token, err := common.GenerateToken(
			common.JwtPayLoad{
				UserID:   1,
				Username: req.Username,
			},
			auth.AccessSecret, auth.AccessExpire,
		)
		if err != nil {
			fmt.Println("生成token失败")
			return "", errors.New("账号或密码错误")
		}
		return token, nil
	} else {
		return "", errors.New("账号或密码错误")
	}
}
