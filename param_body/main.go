package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

//参数：获取body内容

//命令行测试
//$ curl -X POST "http://127.0.0.1:8080/test" -d '{"name":"wang"}'
//返回
//{"name":"wang"}

//命令行测试
//$ curl -X POST "http://127.0.0.1:8080/test" -d 'firstname=wang&lastname=jin'
//结果
//wang,jin,firstname=wang&lastname=jin

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}
		//c.String(http.StatusOK, string(bodyBytes))
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		firstname := c.PostForm("firstname")
		lastname := c.DefaultPostForm("lastname", "default_lastname")
		c.String(http.StatusOK, "%s,%s,%s", firstname, lastname, string(bodyBytes))
	})
	r.Run(":8080")
}

//c.Next()：挂起当前函数，执行后续函数，等后续函数执行完成之后，仍然会回到当前的函数继续执行。在使用的过程中，可以根据其余部分是否继续执行而在 c.Next() 后添加 return 语句。
//c.Abort():中止链的调用。继续执行当前函数，但是阻止了链中后续的函数的调用。
//注意： Abort()和AbortWithStatusJSON` 一直都是中止链的调用，不同之处在于，前者不会返回给前端任何内容，而后者则可以返回给前端一个JSON串。
//参考：https://blog.csdn.net/cyberspecter/article/details/100602552
