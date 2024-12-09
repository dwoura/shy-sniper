package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"market-monitor/api/internal/logic"
	"market-monitor/api/internal/svc"
	"market-monitor/api/internal/types"
)

func SubscribeBinanceAccountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SubscribeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSubscribeBinanceAccountLogic(r.Context(), svcCtx)
		resp, err := l.SubscribeBinanceAccount(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
