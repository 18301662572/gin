package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//验证：结构体验证
//验证规则：https://godoc.org/gopkg.in/go-playground/validator.v9

//命令行测试
//$ curl -X GET "http://127.0.0.1:8080/testing?age=11&name=231&address=dsda"
//返回
//{11 231 dsda}

type Person struct {
	Age     int    `form:"age" binding:"required,gt=10"` //验证规则：如果两个条件同时满足用,表示；如果任意一个满足使用|表示
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(http.StatusInternalServerError, "err:%v", err)
			c.Abort()
			return
		}
		c.String(http.StatusOK, "%v", person)
	})
	r.Run(":8080")
}
