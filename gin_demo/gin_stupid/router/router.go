package router

import (
	"future-go/gin_demo/gin_stupid/consts"
	"future-go/gin_demo/gin_stupid/handler"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Handler 处理器
func retHelloGinAndMethod(c *gin.Context) {
	c.String(http.StatusOK, "hello gin "+strings.ToLower(c.Request.Method)+" method")
}

/* Gin 的路由支持 GET , POST , PUT , DELETE , PATCH , HEAD , OPTIONS 请求，
同时还有一个 Any 函数，可以同时支持以上的所有请求。
*/
func SetupRouter() *gin.Engine {
	r := gin.Default()

	{
		// 添加Get请求路由
		r.GET("/", retHelloGinAndMethod)

		// 添加Post请求路由
		r.POST("/", retHelloGinAndMethod)

		// 添加Put请求路由
		r.PUT("/", retHelloGinAndMethod)

		// 添加Delete请求路由
		r.DELETE("/", retHelloGinAndMethod)

		// 添加Patch请求路由
		r.PATCH("/", retHelloGinAndMethod)

		// 添加Head请求路由
		r.HEAD("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello gin head method")
		})

		// 添加Options请求路由
		r.OPTIONS("/", retHelloGinAndMethod)
	}

	{
		// 添加 user
		// "/:" 该符号表示后面的字符串为一个占位符, 用于将要进行的传值
		// 此时我们的路由为 /user/{name}
		userRouter := r.Group("user")
		{
			userRouter.GET("/:name", handler.UserSave)
			userRouter.GET("", handler.UserSaveByQuery)
			userRouter.POST("", handler.UserSaveByPost)
		}
	}

	// 路由分组( 分组内依旧可以嵌套分组), 提供相同的路由前缀
	{
		v1 := r.Group("/api/v1/group")

		// Any 函数可以通过任何请求
		v1.Any("", retHelloGinAndMethod)
	}

	// 模板tmpl
	{
		/*	LoadHTMLGlob 	将一个目录下所有的模板进行加载
			LoadHTMLFiles	只会加载一个文件，他的参数为可变长参数  */
		r.LoadHTMLGlob(consts.GetTmplPath())

		// 添加静态资源 Bootstrap4, 还可以引入 Jquery,Popper
		r.Static("/statics", "./statics")

		r.GET("/index", handler.Index)
	}

	return r
}
