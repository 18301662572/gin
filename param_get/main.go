package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//参数:获取请求参数

//命令行测试
//$ curl -X GET "http://127.0.0.1:8080/test?firstname=wang&lastname=22"
//返回
//wang,22

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		firstname := c.Query("firstname")
		lastname := c.DefaultQuery("lastname", "")
		c.String(http.StatusOK, "%s,%s", firstname, lastname)
	})
	r.Run(":8080")
}
