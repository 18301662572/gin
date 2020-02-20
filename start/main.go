package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//初识 gin

//命令行测试
//curl -X GET "http://127.0.0.1:8080/ping"

func main(){
	r:=gin.Default()
	r.GET("/ping",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"msg":"pong",
		})
	})
	r.Run(":8080")//默认是8080
}
