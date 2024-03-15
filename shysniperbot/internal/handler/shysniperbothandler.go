package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"shy-sniper/shysniperbot/internal/logic"
	"shy-sniper/shysniperbot/internal/svc"
	"shy-sniper/shysniperbot/internal/types"
)

func ShysniperbotHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewShysniperbotLogic(r.Context(), svcCtx)
		resp, err := l.Shysniperbot(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
