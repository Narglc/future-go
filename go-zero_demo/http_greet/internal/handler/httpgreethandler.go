package handler

import (
	"net/http"

	"future-go/go-zero_demo/http_greet/internal/logic"
	"future-go/go-zero_demo/http_greet/internal/svc"
	"future-go/go-zero_demo/http_greet/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Http_greetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewHttp_greetLogic(r.Context(), svcCtx)
		resp, err := l.Http_greet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
