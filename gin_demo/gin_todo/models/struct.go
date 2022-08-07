package models

import (
	"gorm.io/gorm"
)

type (

	// 定义原始的数据库字段
	TodoModel struct {
		gorm.Model        // 把 ID,CreatedAt,UpdatedAt 和 DeletedAt 这四个字段嵌入到我们定义好的 todoModel 结构体中，一般数据表中都会用到这四个字段。
		Title      string `json:"title"`
		Completed  int    `json:"completed"`
	}

	// 处理返回的字段, 精简gorm.Model, 只留下ID节点, 用于序列化
	TransformedTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)
