package router

import (
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

	return r
}
