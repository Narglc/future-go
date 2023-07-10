package handler

import (
	"future-go/gin_demo/gin_stupid/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
参数获取方法:
	- Param
	- Query
	- DefaultQuery
	- QueryArray
	- QueryMap
	- PostForm
	- DefaultPostForm
*/

// 保存 User
func UserSave(c *gin.Context) {
	// GET方法 -> c.Param 可以获取路由路径中的参数
	username := c.Param("name")
	c.String(http.StatusOK, "用户:"+username+" 已经保存")
}

// 通过 query 方法获取参数
func UserSaveByQuery(c *gin.Context) {
	// GET方法 -> c.Query 获取参数
	username := c.Query("name")

	// GET方法 -> c.DefaultQuery 提供默认值
	age := c.DefaultQuery("age", "20")
	c.JSON(http.StatusOK, gin.H{
		"name": username,
		"age":  age,
	})
}

func UserSaveByPost(c *gin.Context) {
	// POST方法 -> c.PostForm
	username := c.PostForm("name")
	age := c.DefaultPostForm("age", "18")

	c.JSON(http.StatusOK, gin.H{
		"name": username,
		"age":  age,
	})
}

func UserRegister(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		log.Println("err ->", err.Error())
		context.String(http.StatusBadRequest, "输入的数据不合法")
	}
	log.Println("email", user.Email, "password", user.Password, "password again", user.PasswordAgain)
	context.Redirect(http.StatusMovedPermanently, "/")
}
