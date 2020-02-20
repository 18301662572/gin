package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//路由：泛绑定

//命令行测试
//$ curl -X GET "http://127.0.0.1:8080/user/action"
//$ curl -X GET "http://127.0.0.1:8080/user/11"
//$ curl -X GET "http://127.0.0.1:8080/user/22"
//返回都是
//hello world

func main(){
	r:=gin.Default()
	//GET 泛绑定前缀
	r.GET("/user/*action", func(c *gin.Context) {
		c.String(http.StatusOK,"hello world")
	})
	r.Run(":8080")
}
