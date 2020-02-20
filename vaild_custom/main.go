package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/gin-gonic/gin.v1/binding"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"reflect"
	"time"
)

//验证：自定义验证
//验证规则：https://godoc.org/gopkg.in/go-playground/validator.v8

//自定义验证
func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Unix() < date.Unix() {
			return false
		}
	}
	return true
}

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required" validate:"bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()

	//注册自定义规则
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	r.GET("/bookable", func(c *gin.Context) {
		var book Booking
		if err := c.ShouldBind(&book); err != nil {
			c.String(http.StatusInternalServerError, "err:%v", err)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "ok",
			"book": book,
		})
	})
	r.Run(":8080")
}
