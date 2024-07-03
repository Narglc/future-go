package svc

import (
	"future-go/go-zero_demo/micro_mall/user/database"
	"future-go/go-zero_demo/micro_mall/user/internal/config"
	"future-go/go-zero_demo/micro_mall/user/internal/dao"
	"future-go/go-zero_demo/micro_mall/user/internal/repo"
)

type ServiceContext struct {
	Config   config.Config
	UserRepo repo.UserRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	connect := database.Connect(c.Mysql.DataSource, c.CacheRedis)
	return &ServiceContext{
		Config:   c,
		UserRepo: dao.NewUserDao(connect),
	}
}
