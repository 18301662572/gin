package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//其他补充：优雅关停服务器

//1.命令行测试
//$ curl -X GET "http://127.0.0.1:8085/test"
//2.立即关停服务器
//返回结果
//1.服务器返回：
// 2020/02/20 20:27:11 shutdown server...
//2020/02/20 - 20:27:14 | 200 |   10.0005748s |       127.0.0.1 | GET      /test
// 2020/02/20 20:27:14 server exiting
//2.客户端返回
//hello test

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.String(200, "hello test")
	})
	srv := http.Server{
		Addr:    ":8085",
		Handler: r,
	}
	//把server放到一个协程里面
	go func() {
		if err := srv.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatalf("listen err:%s\n", err)
		}
	}()
	//创建一个接受阻塞信号的通道
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) //os.Signal退出的信号捕获
	<-quit                                               //阻塞channel
	log.Println("shutdown server...")
	//创建超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Println("server exiting")
}
