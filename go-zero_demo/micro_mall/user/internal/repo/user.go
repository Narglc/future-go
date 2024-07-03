package repo

import (
	"context"
	"future-go/go-zero_demo/micro_mall/user/internal/model"
)

type UserRepo interface {
	Save(ctx context.Context, user *model.User) error
	FindById(ctx context.Context, id int64) (user *model.User, err error)
}
