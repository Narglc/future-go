package logic

import (
	"context"

	"future-go/go-zero_demo/micro_mall/gate/internal/svc"
	"future-go/go-zero_demo/micro_mall/gate/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GateTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGateTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GateTestLogic {
	return &GateTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GateTestLogic) GateTest(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	logx.Infof("执行test业务处理")
	return
}
