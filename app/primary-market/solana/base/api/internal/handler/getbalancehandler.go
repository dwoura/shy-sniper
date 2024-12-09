package handler

import (
	"base/api/internal/types"
	"net/http"

	"base/api/internal/logic"
	"base/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getBalanceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetBalanceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetBalanceLogic(r.Context(), svcCtx)
		resp, err := l.GetBalance(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
