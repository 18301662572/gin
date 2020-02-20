package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

//中间件：gin中间件

//命令行测试
//$ curl -X GET "http://127.0.0.1:8080/test?name=dsda"
//返回
//dsda

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_name")
		panic("test panic")
		c.String(200, "%s", name)
	})
	r.Run()
}
