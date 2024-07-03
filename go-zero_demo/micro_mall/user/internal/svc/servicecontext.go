package svc

import (
	database "command-line-arguments/Users/narglc/go/src/future-go/go-zero_demo/micro_mall/user/database/sqlx.go"
	"command-line-arguments/Users/narglc/go/src/future-go/go-zero_demo/micro_mall/user/internal/dao"
	"command-line-arguments/Users/narglc/go/src/future-go/go-zero_demo/micro_mall/user/internal/repo"
	"future-go/go-zero_demo/micro_mall/user/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	UserRepo repo.UserRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	connect := database.Connect(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		UserRepo: dao.NewUserDao(connect),
	}
}
