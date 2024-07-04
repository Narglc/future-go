package logic

import (
	"context"
	"fmt"
	"time"

	"future-go/go-zero_demo/micro_mall/gate/internal/svc"
	"future-go/go-zero_demo/micro_mall/gate/internal/types"
	"future-go/go-zero_demo/micro_mall/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) GetUser(req *types.IdRequest) (resp *types.UserResponse, err error) {
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

func (l *UserLogic) Login(req *types.UserLoginReq) (resp *types.UserLoginRsp, err error) {
	// todo: add your logic here and delete this line

	return
}

func (l *UserLogic) Register(req *types.UserRegisterReq) (resp *types.UserRegisterRsp, err error) {
	// todo: add your logic here and delete this line

	// 超时上下文控制,10ms会超时
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer func() {
		fmt.Printf("5s超时了")
		cancelFunc()
	}()

	// 根据用户id去user服务获取用户信息
	rsp, err := l.svcCtx.UserRpc.SaveUser(ctx, &user.UserRequest{
		Name:   req.Name,
		Gender: req.Gender,
	})

	if err != nil {
		return nil, err
	}

	return &types.UserRegisterRsp{
		Id:      rsp.Id,
		Name:    rsp.Name,
		Gender:  rsp.Gender,
		Message: "succ",
	}, nil
}
