package logic

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"future-go/go-zero_demo/micro_mall/user/internal/model"
	"future-go/go-zero_demo/micro_mall/user/internal/svc"
	"future-go/go-zero_demo/micro_mall/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (u *UserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	// 修改为数据库、redis实现
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	//userId, _ := strconv.Atoi(in.GetId())
	userId, _ := strconv.ParseInt(in.GetId(), 10, 64)

	userData, err := u.svcCtx.UserRepo.FindById(ctx, int64(userId))

	if err != nil {
		return nil, err
	}

	return &user.UserResponse{
		Id:     in.GetId(),
		Name:   userData.Name,
		Gender: userData.Gender,
	}, nil
}

func (u *UserLogic) SaveUser(in *user.UserRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	data := &model.User{
		Name:   in.Name,
		Gender: in.Gender,
	}

	err := u.svcCtx.UserRepo.Save(context.Background(), data)

	if err != nil {
		return nil, err
	}

	userId := fmt.Sprintf("%d", data.Id)
	return &user.UserResponse{
		Id:     userId, //strconv.FormatInt(data.Id, 10),
		Name:   data.Name,
		Gender: data.Gender,
	}, nil
}
