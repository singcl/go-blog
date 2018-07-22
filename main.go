package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 默认是gin是debug模式
	// 这里设置以release模式启动
	gin.SetMode(gin.ReleaseMode)

	app := gin.New()

	// 设置日志文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 使用日志中间件
	app.Use(gin.Logger())

	// 设置静态资源文件
	app.Static("/static", "./static")
	app.Run(":8000")
}
