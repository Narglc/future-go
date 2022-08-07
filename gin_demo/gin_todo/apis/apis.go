package apis

import (
	"fmt"
	"net/http"
	"strconv"

	. "future-go/gin_demo/gin_todo/dao"
	. "future-go/gin_demo/gin_todo/models"

	"github.com/gin-gonic/gin"
)

// createTodo 方法添加一条新的 todo 数据
// 使用 ApiFox POST方法+Body数据请求
func CreateTodo(c *gin.Context) {
	title := c.PostForm("title")
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := TodoModel{
		Title:     title,
		Completed: completed,
	}

	Db.Save(&todo)

	c.JSON(http.StatusCreated, gin.H{ // gin.H = map[string]interface{}
		"status":     http.StatusCreated,
		"message":    "Todo item created successfully!",
		"resourceId": todo.ID,
	})
}

// fetchSingleTodo方法返回一条 todo 数据
func FetchAllTodo(c *gin.Context) {
	var todos []TodoModel
	var _todos []TransformedTodo

	Db.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No todo list!",
		})
		return
	}

	// 转化 todos 数据，用于格式化
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		}
		_todos = append(_todos, TransformedTodo{
			ID:        item.ID,
			Title:     item.Title,
			Completed: completed,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   _todos,
	})
}

// fetchSingleTodo方法返回一条 todo 数据
func FetchSingleTodo(c *gin.Context) {
	var todo TodoModel
	todoId := c.Param("id")

	Db.First(&todo, todoId)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No todo list!",
		})
		return
	}

	completed := false
	if todo.Completed == 1 {
		completed = true
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": TransformedTodo{
			ID:        todo.ID,
			Title:     todo.Title,
			Completed: completed,
		},
	})

}

// updateTodo 方法 更新 todo 数据
func UpdateTodo(c *gin.Context) {
	var todo TodoModel
	todoID := c.Param("id")

	Db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No todo list!",
		})
		return
	}

	// Db.Model() 指定到要操作的那行记录
	Db.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	Db.Model(&todo).Update("completed", completed)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo updated successfully!",
	})
}

// deleteTodo 方法依据 id 删除一条todo 数据
func DeleteTodo(c *gin.Context) {
	var todo TodoModel
	todoId := c.Param("id")
	fmt.Printf("--------->id: %v", todoId)

	Db.First(&todo, todoId)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No todo list!",
		})
		return
	}

	Db.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo deteled successfully!",
	})
}
