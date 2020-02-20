package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//路由：多种请求类型(请求类型区分大小写)

//命令行测试
//curl -X GET "http://127.0.0.1:8080/get"

func main(){
	r:=gin.Default()
	//GET、POST ("路由","回调方法")
	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK,"get")
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK,"post")
	})
	//Handle ("请求方式","请求路由","回调方法")
	r.Handle("DELETE","/delete", func(c *gin.Context) {
		c.String(http.StatusOK,"delete")
	})
	//Any 创建8种请求类型
	r.Any("/any", func(c *gin.Context) {
		c.String(http.StatusOK,"any")
	})
	r.Run(":8080")
}
