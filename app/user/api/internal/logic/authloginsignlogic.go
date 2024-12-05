package logic

import (
	"context"
	"errors"
	"user/api/internal/svc"
	"user/api/internal/types"
	"user/common"
	"user/entity"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spruceid/siwe-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLoginSignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthLoginSignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLoginSignLogic {
	return &AuthLoginSignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLoginSignLogic) AuthLoginSign(req *types.LoginSignRequest) (resp string, err error) {
	// 验签
	message, err := siwe.ParseMessage(req.Message)
	if err != nil {
		return "", err
	}
	getNonce, err := l.svcCtx.Redis.Get(message.GetNonce()) // 检查 nonce 是否存在
	isValid, err := message.ValidNow()
	if err != nil {
		return "", err
	}
	// 判断nonce 还有时间有没有过期
	if getNonce != "true" || isValid != true {
		return "", errors.New("wrong message")
	}
	publicKey, err := message.VerifyEIP191(req.Signature) // 验证签名和nonce
	if err != nil {
		return "", err
	}
	userAddr := crypto.PubkeyToAddress(*publicKey)

	DB := l.svcCtx.DB
	// 查询用户，没有则为其注册，怎么防止注册攻击，预检？
	result := DB.Table("users").Where("address = ?", userAddr.String()).FirstOrCreate(&entity.Users{
		Address: userAddr.String(),
	}) // &user 为存放数据的载体
	if result.Error != nil {
		return "", result.Error
	}

	// 验证成功，生成并返回 jwt
	auth := l.svcCtx.Config.Auth
	token, err := common.GenerateToken(
		common.JwtPayLoad{
			Address: userAddr.String(),
		},
		auth.AccessSecret, auth.AccessExpire,
	)
	if err != nil {
		//fmt.Println("生成token失败")
		return "", errors.New("failed to generate jwt token")
	}
	return token, nil
}
