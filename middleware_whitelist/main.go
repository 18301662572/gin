package main

import "github.com/gin-gonic/gin"

//中间件：自定义中间件

//命令行测试
//$ curl -X GET "http://127.0.0.1:8080/test"
//返回
//127.0.0.1 , not in ipList

func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{
			"127.0.0.2",
		}
		flag := false
		cliantIP := c.ClientIP()
		for _, host := range ipList {
			if cliantIP == host {
				flag = true
				break
			}
		}
		if !flag {
			c.String(401, "%s , not in ipList", cliantIP)
			c.Abort()
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(IPAuthMiddleware())
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello test")
	})
	r.Run()
}
