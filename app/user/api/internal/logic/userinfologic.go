package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"user/entity"

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
	// 从请求头中获取token,解析出来
	userAddr := l.ctx.Value("address").(string) // jwt 中间件已经将 payload 解析到 ctx 中
	fmt.Println(userAddr)

	// 结合 gorm
	var user entity.Users
	tx := l.svcCtx.DB.Table("users").Where("address = ?", userAddr).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	data, _ := json.Marshal(user)
	_ = json.Unmarshal(data, &resp)
	return resp, nil
}
