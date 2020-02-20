package main

import (
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"net/http"
)

//验证：多语言验证

type Person struct {
	Age     int    `form:"age" validate:"required,gt=10"` //v9验证使用validate，而不是binding
	Name    string `form:"name" validate:"required"`
	Address string `form:"address" validate:"required"`
}

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

func main() {
	Validate = validator.New()
	zh := zh2.New()
	en := en2.New()
	Uni := ut.New(zh, en) //翻译器（支持的语言）

	r := gin.Default()
	r.GET("/testing", func(c *gin.Context) {
		locale := c.DefaultQuery("locate", "zh")
		trans, _ := Uni.GetTranslator(locale)
		switch locale {
		case "zh":
			zh_translations.RegisterDefaultTranslations(Validate, trans)
		case "en":
			en_translations.RegisterDefaultTranslations(Validate, trans)
		default:
			zh_translations.RegisterDefaultTranslations(Validate, trans)
		}

		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(http.StatusInternalServerError, "err:%v", err)
			c.Abort()
			return
		}
		//验证结构体
		if err := Validate.Struct(person); err != nil {
			errs := err.(validator.ValidationErrors)
			sliceErrs := []string{}
			for _, e := range errs {
				sliceErrs = append(sliceErrs, e.Translate(trans))
			}
			c.String(http.StatusInternalServerError, "%v", sliceErrs)
			c.Abort()
			return
		}
		c.String(200, "%v", person)
	})
	r.Run("")
}
