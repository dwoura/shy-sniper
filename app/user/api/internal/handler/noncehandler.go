package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/api/internal/logic"
	"user/api/internal/svc"
)

func nonceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewNonceLogic(r.Context(), svcCtx)
		resp, err := l.Nonce()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
