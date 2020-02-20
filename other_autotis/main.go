package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

//其他补充：自动创建证书

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello test")
	})

	//1.生成一个本地的密钥
	//2.把密钥发给证书颁发机构，获取一个私钥
	//3.私钥验证成功，把私钥信息保存起来，下次请求使用私钥进行加密，达到一个自动下载证书的功能
	autotls.Run(r, "www.itpp.tk")
}
