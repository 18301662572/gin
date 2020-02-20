package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//路由：参数作为URL传递

//命令行测试
//$ curl -X GET "http://127.0.0.1:8080/wangjin/32"
//返回
// {"id":"32","name":"wangjin"}

func main(){
	r:=gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"name":c.Param("name"),
			"id":c.Param("id"),
		})
	})
	r.Run(":8080")
}
