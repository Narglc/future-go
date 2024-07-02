package handler

import (
	"net/http"

	"future-go/go-zero_demo/micro_mall/order/internal/logic"
	"future-go/go-zero_demo/micro_mall/order/internal/svc"
	"future-go/go-zero_demo/micro_mall/order/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOrderLogic(r.Context(), svcCtx)
		resp, err := l.Order(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
