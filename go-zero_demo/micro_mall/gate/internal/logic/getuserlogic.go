package logic

import (
	"context"
	"fmt"

	"future-go/go-zero_demo/micro_mall/gate/internal/svc"
	"future-go/go-zero_demo/micro_mall/gate/internal/types"
	"future-go/go-zero_demo/micro_mall/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.IdRequest) (resp *types.UserResponse, err error) {
	// todo: add your logic here and delete this line
	userId := req.Id

	// 根据用户id去user服务获取用户信息
	userResponse, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user.IdRequest{
		Id: userId,
	})

	if err != nil {
		return nil, err
	}

	resp = &types.UserResponse{
		Id:       userResponse.Id,
		Name:     userResponse.Name,
		UserName: fmt.Sprintf("gender:%s", userResponse.Gender),
	}

	return
}
