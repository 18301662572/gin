package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//参数：参数绑定到结构体

//命令行测试
//$ curl -X GET "http://127.0.0.1:8080/testting?name=wang&address=wuhan&birthday=2009-07-09"
//$ curl -X POST "http://127.0.0.1:8080/testting" -d 'name=wang&birthday=2009-07-09'
//$ curl -H "Content-Type:application/json" -X POST "http://127.0.0.1:8080/testting" -d '{"name":"wang","birthday":"2009-07-09"}'

//返回
//{wang  2009-07-09 00:00:00 +0800 CST}

//结构体
type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	r.POST("/testting", testting)
	r.GET("/testting", testting)
	r.Run(":8080")
}

func testting(c *gin.Context) {
	var person Person
	//根据请求的content-type 来做不同的binding操作
	if err := c.ShouldBind(&person); err == nil {
		c.String(http.StatusOK, "%v", person)
	} else {
		c.String(http.StatusOK, "person bind error:%v", err)
	}
}
