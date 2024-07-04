package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserMiddleware struct {
}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (m *UserMiddleware) RegAndLoginHandle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		logx.Info("login 和 register 前面执行")

		// Passthrough to next handler if need
		next(w, r)

		logx.Info("login 和 register 后面面执行")

	}
}

func (m *UserMiddleware) GlobalHandle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		logx.Info("global 全局前面执行")

		// Passthrough to next handler if need
		next(w, r)

		logx.Info("global 全局后面面执行")

	}
}
