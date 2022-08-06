package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	// 创建一个数据库的连接
	// parseTime: 正确处理 time.Time
	dsn := "root:MySql1990328@tcp(127.0.0.1:3306)/todo?charset=utf8&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // 此处不能使用 := , 会将Db视为局部值
	if err != nil {
		panic("fail to connect database.")
	}

	// 迁移 schema (Gorm 有迁移工具 ，在调用 ‘init’ 函数初始化的时候已经初始化了。当我们运行应用程序时，它将创建一个连接然后进行迁移。)
	err = Db.AutoMigrate(&TodoModel{})
	if err != nil {
		panic(err)
	}
}
