package global

import (
	"log"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

/*
 * @MethodName
 * @Description 转中文错误
 * @Author khr
 * @Date 2023/4/23 9:22
 */

var trans ut.Translator
var Validate *validator.Validate

var uni *ut.UniversalTranslator

// 参数验证，转中文
func init() {
	//修改gin框架中的validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)

		trans, ok = uni.GetTranslator(LANGUAGE)
		if !ok {
			panic("参数校验语言翻译失败")
			//return fmt.Errorf("uni.GetTranslator(%s)", LANGUAGE)
		}
		switch LANGUAGE {
		case "en":
			en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			en_translations.RegisterDefaultTranslations(v, trans)
		}
	}

	log.Printf("翻译初始化成功")

}

// Translate 翻译错误信息
func Translate(err error) map[string][]string {
	var result = make(map[string][]string)
	errors := err.(validator.ValidationErrors)
	for _, err := range errors {
		result[err.Field()] = append(result[err.Field()], err.Translate(trans))
	}
	return result
}
func TransErrToZH(err error) string {
	var result string
	errors := err.(validator.ValidationErrors)
	for _, err := range errors {
		result = err.Translate(trans)
	}
	return result
}
