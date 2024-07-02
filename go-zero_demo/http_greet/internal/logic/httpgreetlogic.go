package logic

import (
	"context"

	"future-go/go-zero_demo/http_greet/internal/svc"
	"future-go/go-zero_demo/http_greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Http_greetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHttp_greetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Http_greetLogic {
	return &Http_greetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Http_greetLogic) Http_greet(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
