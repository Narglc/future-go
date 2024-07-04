package logic

import (
	"context"
	"fmt"
	"time"

	"future-go/go-zero_demo/micro_mall/gate/internal/errorx"
	"future-go/go-zero_demo/micro_mall/gate/internal/svc"
	"future-go/go-zero_demo/micro_mall/gate/internal/types"
	"future-go/go-zero_demo/micro_mall/user/types/user"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logc"
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

	// 认证通过后，可从token中获取用于id userid， 框架会将其保存在 ctx 上下文中
	userId := l.ctx.Value("userId")
	// logc 相对 logx 增加了 trace 信息，使用上第一个参数要传入 ctx
	logc.Infof(l.ctx, "获取的token内容:%s \n", userId)

	// 自定义错误
	if req.Id == "1" {
		return nil, errorx.ParamsError
	}

	// 根据用户id去user服务获取用户信息
	userResponse, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user.IdRequest{
		Id: req.Id,
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

func (l *UserLogic) getToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *UserLogic) Login(req *types.UserLoginReq) (resp *types.UserLoginRsp, err error) {
	// todo: add your logic here and delete this line
	logx.Infof("执行 Login 业务处理")

	userId := 100
	auth := l.svcCtx.Config.Auth
	token, err := l.getToken(auth.AccessSecret, time.Now().Unix(), auth.AccessExpire, int64(userId))
	if err != nil {
		return nil, err
	}
	resp = &types.UserLoginRsp{
		Token: token,
	}
	return
}

func (l *UserLogic) Register(req *types.UserRegisterReq) (resp *types.UserRegisterRsp, err error) {
	// todo: add your logic here and delete this line
	logx.Infof("执行 Register 业务处理")

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
