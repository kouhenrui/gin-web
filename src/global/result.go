package global

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Ctx *gin.Context
}

// 返回的结果：
type ResultCont struct {
	Code int         `json:"code"` //提示代码
	Msg  interface{} `json:"msg"`  //提示信息
	Data interface{} `json:"data"` //数据
}

func NewResult(ctx *gin.Context) *Result {
	return &Result{Ctx: ctx}
}

// 成功
func (r *Result) Success(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := ResultCont{}
	res.Code = http.StatusOK
	res.Msg = ""
	res.Data = data
	r.Ctx.JSON(http.StatusOK, res)
}

func (r *Result) Succ() {
	res := ResultCont{}
	res.Code = http.StatusOK
	res.Msg = "success"
	res.Data = ""
	r.Ctx.JSON(http.StatusOK, res)
}

// 出错
func (r *Result) Error(code int, msg interface{}) {
	res := ResultCont{}
	res.Code = code
	res.Msg = msg
	res.Data = gin.H{}
	r.Ctx.JSON(http.StatusOK, res)
}
func (r *Result) Err(msg interface{}) {
	res := ResultCont{}
	res.Code = 700
	res.Msg = msg
	res.Data = gin.H{}
	fmt.Println("mag", res.Msg)
	r.Ctx.JSON(res.Code, res)
}

//func (r *Result) DiyErr(code int, msg interface{}) {
//	if msg == nil {
//		msg = gin.H{}
//	}
//	res := ResultCont{}
//	res.Code = code
//	res.Msg = msg
//	res.Data = ""
//	r.Ctx.JSON(code, res)
//}
