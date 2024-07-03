package dao

import (
	database "command-line-arguments/Users/narglc/go/src/future-go/go-zero_demo/micro_mall/user/database/sqlx.go"
	"context"
	"future-go/go-zero_demo/micro_mall/user/internal/model"
)

type userDao struct {
	conn *database.DBConn
}

func NewUserDao(conn *database.DBConn) *userDao {
	return &userDao{
		conn: conn,
	}
}

func (d *userDao) Save(ctx context.Context, user *model.User) error {
	// TODO: 实际的代码实现处, 数据库连接使用 d.conn.Conn

	return nil
}
