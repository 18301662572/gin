package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//路由： 静态文件夹绑定

//因为静态文件夹找不到，所以需要命令行开启服务
//1.进入这个文件夹
//2.运行 go build -o router_static && ./router_static
//打开 gitbush
//jin@DESKTOP-6O70DRL MINGW64 /g/gowork/src/github.com/18301662572/gin_test_project/router_static
//$ go build -o router_static && ./router_static

//3.客户端测试
//curl -X GET "http://127.0.0.1:8080/static/b.html"

func main(){
	r:=gin.Default()
	//静态文件夹绑定
	//Static("路由"，"文件夹绝对路径")
	r.Static("/assets","./assets")
	r.StaticFS("/static",http.Dir("static"))
	r.StaticFile("/favicon.ico","./favicon.ico")
	r.Run(":8080")
}