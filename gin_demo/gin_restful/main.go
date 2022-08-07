package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

/* RestFul API 简介
简单地说，就是
- 用GET请求来查询获取数据；
- 用POST请求来创建数据；
- 用PUT请求来更新数据；
- 用DELETE请求来删除数据
*/

//函数会返回状态码是200，JSON格式的数据key是message，value是"Hello Golang"
// http.StatusOK == 200
func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Golang~",
	})
}

func main() {
	// 全局设置环境
	//gin.SetMode(gin.DebugMode) // gin.ReleaseMode

	// 获得路由实例(创建默认的路由引擎)
	router := gin.Default()

	// 用户发送GET请求的时候，会触发sayHello函数
	router.GET("/hello", sayHello)

	// GET/POST/PUT/DELETE四种方法简单测试
	{
		// GET http://localhost:8080/test?firstname=wyf&lastname=123
		router.GET("/test", func(c *gin.Context) {
			firstname := c.DefaultQuery("firstname", "Guest")
			lastname := c.Query("lastname")

			c.JSON(200, gin.H{
				"method":  "GET",
				"message": fmt.Sprintf("Hello,%s %s", firstname, lastname),
			})
		})

		// POST http://localhost:8080/test
		// Body 传参message、nick
		// POST 也可与 GET 方法混合使用
		router.POST("/test", func(c *gin.Context) {
			message := c.PostForm("message")
			nick := c.DefaultPostForm("nick", "anonymous")

			c.String(200, "method:%s\nmessage:%s\nnick:%s\n", "POST", message, nick)
		})

		router.PUT("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"method": "PUT",
			})
		})

		router.DELETE("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"method": "DELETE",
			})
		})
	}

	v1 := router.Group("/api/v1/userinfo")
	// 创建五个routes
	{
		v1.POST("/", CreateUser)
		v1.GET("/", FetchAllUsers)
		v1.GET("/:id", FetchSingleUser)
		v1.PUT("/:id", UpdateUser)
		v1.DELETE("/:id", DeleteUser)
	}

	// 默认在localhost:8080端口启动服务,也可手动指定 router.Run(":9090")
	router.Run()
}
