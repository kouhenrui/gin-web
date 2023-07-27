package middleWare

import (
	"fmt"
	"gin-web/src/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

// 返回的结果：
type Result struct {
	Code int         `json:"code"` //提示代码
	Msg  string      `json:"msg"`  //提示信息
	Data interface{} `json:"data"` //数据
}
type AuthErr struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func (a *AuthErr) AuthErr() *AuthErr {
	return a
}

func Recover(c *gin.Context) {
	defer func() {

		result := global.NewResult(c)
		if r := recover(); r != nil {
			fmt.Println("打印错误信息:", r)
			//打印错误堆栈信息
			fmt.Printf("panic: %v\n", r)

			debug.PrintStack()
			result.Error(http.StatusInternalServerError, errorToString(r))
			return
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()

		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	fmt.Println(r)
	switch v := r.(type) {
	//fmt.Println("错误类型：",v)
	case error:
		return v.Error()

	case *AuthErr:
		return r.(string)
	default:
		return r.(string)
	}
}
