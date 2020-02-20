package main

import "github.com/gin-gonic/gin"

//其他补充：模板渲染
//详细的功能演示，参照官网： https://github.com/gin-gonic/gin
//gin中文翻译网：https://www.kancloud.cn/shuangdeyu/gin_book/949411

//启动服务：因为有模板，所在的目录是不一致的，所以需要在命令行手动启动服务
//1.右键进入cmd
//2.运行 go run main.go
//访问：$ curl -X GET "HTTP://127.0.0.1:8080/a"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.GET("/a", func(c *gin.Context) {
		c.HTML(200, "a.html", gin.H{
			"Title": "a.html",
		})
	})
	r.Run()
}
