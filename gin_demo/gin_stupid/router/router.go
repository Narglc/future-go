package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 添加Get请求路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin")
	})
	return r
}
