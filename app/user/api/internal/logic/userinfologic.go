package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	// 从请求头中获取token,解析出来
	userId := l.ctx.Value("userId").(json.Number)
	fmt.Println(userId)
	fmt.Printf("数据类型:%v,%T\n", userId, userId)
	username := l.ctx.Value("username").(string)
	fmt.Println(username)
	uid, _ := userId.Int64()
	// 结合 gorm
	return &types.UserInfoResponse{
		Id:       uid,
		Username: username,
	}, nil
}
