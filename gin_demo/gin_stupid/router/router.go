package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* Gin 的路由支持 GET , POST , PUT , DELETE , PATCH , HEAD , OPTIONS 请求，
同时还有一个 Any 函数，可以同时支持以上的所有请求。
*/
func SetupRouter() *gin.Engine {
	r := gin.Default()

	{
		// 添加Get请求路由
		r.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello gin get method")
		})

		// 添加Post请求路由
		r.POST("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello gin post method")
		})

		// 添加Put请求路由
		r.PUT("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello gin put method")
		})

		// 添加Delete请求路由
		r.DELETE("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello gin delete method")
		})

		// 添加Patch请求路由
		r.PATCH("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello gin patch method")
		})

		// 添加Head请求路由
		r.HEAD("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello gin head method")
		})

		// 添加Options请求路由
		r.OPTIONS("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello gin options method")
		})

	}

	return r
}
