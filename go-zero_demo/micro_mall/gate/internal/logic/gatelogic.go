package logic

import (
	"context"

	"future-go/go-zero_demo/micro_mall/gate/internal/svc"
	"future-go/go-zero_demo/micro_mall/gate/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GateLogic {
	return &GateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GateLogic) Gate(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = &types.Response{
		Message: "you got it!",
	}
	return
}
