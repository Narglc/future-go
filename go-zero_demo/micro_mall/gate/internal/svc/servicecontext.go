package svc

import (
	"future-go/go-zero_demo/micro_mall/gate/internal/config"
	"future-go/go-zero_demo/micro_mall/gate/internal/middleware"
	"future-go/go-zero_demo/micro_mall/user/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	UserRpc        userclient.User
	UserMiddleware *middleware.UserMiddleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		UserMiddleware: middleware.NewUserMiddleware(),
	}
}
