package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

//gin 框架 文件上传（单个文件/多个文件）

//使用cmd 命令上传单个文件
//curl -X POST http://localhost:8080/uploadSingle \
//-F "file=@/C:/Users/jin/Desktop/go/aa.txt" \
//-H "Content-Type: multipart/form-data"

//使用cmd 命令上传多个文件
//curl -X POST http://localhost:8080/uploadMul \
//-F "files=@/Users/appleboy/test1.zip" \
//-F "files=@/Users/appleboy/test2.zip" \
//-H "Content-Type: multipart/form-data"

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*")

	router.GET("/uploadMul", func(c *gin.Context) {
		c.HTML(200, "multipleupload.html", gin.H{})
	})

	router.GET("/uploadSingle", func(c *gin.Context) {
		c.HTML(200, "singleupload.html", gin.H{})
	})

	// 给表单限制上传大小 (默认 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	//单文件上传
	router.POST("/uploadSingle", SingleUploadHandler)

	//多文件上传
	router.POST("/uploadMul", MultipleUploadHandler)

	//路由
	router.Run(":8080")
}

func SingleUploadHandler(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	filename := "file/" + filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email))
}

//多文件上传
func MultipleUploadHandler(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["files"]
	for _, file := range files {
		filename := "file/" + filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files with fields name=%s and email=%s.", len(files), name, email))
}
