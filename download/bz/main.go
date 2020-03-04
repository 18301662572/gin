package bz

import (
	"io"
	"net/http"
	"os"
)

//直接下载文件

func main() {
	url := "https://www.baidu.com/img/bd_logo1.png"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//创建文件目录
	f, err := os.Create("file/bd_logo.png")
	if err != nil {
		panic(err)
	}
	//拷贝
	io.Copy(f, res.Body)
}
