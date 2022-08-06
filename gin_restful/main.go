package main

import "github.com/gin-gonic/gin"

func main() {
	// 获得路由实例
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
