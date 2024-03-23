package logic

import (
	"base/rpc/base"
	"base/rpc/internal/svc"
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"math/big"
)

type GetBalanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBalanceLogic {
	return &GetBalanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBalanceLogic) GetBalance(in *base.GetBalanceReq) (*base.GetBalanceResp, error) {
	pubKey := solana.MustPublicKeyFromBase58(in.Address)
	out, err := l.svcCtx.SolanaClient.GetBalance(context.TODO(), pubKey, rpc.CommitmentFinalized)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	spew.Dump(out)
	spew.Dump(out.Value) // total lamports on the account; 1 sol = 1000000000 lamports

	var lamportsOnAccount = new(big.Float).SetUint64(uint64(out.Value))
	// Convert lamports to sol:
	var solBalance = new(big.Float).Quo(lamportsOnAccount, new(big.Float).SetUint64(solana.LAMPORTS_PER_SOL))
	return &base.GetBalanceResp{
		Balance: solBalance.Text('f', 10),
	}, nil
}
