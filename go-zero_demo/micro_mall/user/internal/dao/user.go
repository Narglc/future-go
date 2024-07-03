package dao

import (
	"context"
	"fmt"
	"future-go/go-zero_demo/micro_mall/user/database"
	"future-go/go-zero_demo/micro_mall/user/internal/model"
)

type userDao struct {
	*database.DBConn
}

func NewUserDao(conn *database.DBConn) *userDao {
	return &userDao{
		conn,
	}
}

func (d *userDao) Save(ctx context.Context, user *model.User) error {
	// 实际的代码实现处, 数据库连接使用 d.Conn
	sql := fmt.Sprintf("insert into %s (name, gender) values (?,?)", user.TableName())
	result, err := d.Conn.ExecCtx(ctx, sql, user.Name, user.Gender)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = id
	return nil
}