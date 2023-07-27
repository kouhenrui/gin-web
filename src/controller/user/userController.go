package user

import (
	"fmt"
	"gin-web/src/dto/reqDto"
	"gin-web/src/global"
	"gin-web/src/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Routers(e *gin.Engine) {

	userGroup := e.Group("/api/user")
	{
		userGroup.POST("/login", newLogin)
		userGroup.POST("/sign", signUser)
		userGroup.GET("/info", getUserInfo)
		userGroup.POST("/post/message", postMessage)
		userGroup.PUT("/put/user", updateUser)
		//userGroup.POST("/register", rejisterUser)
	}

}

/*
 * @MethodName newLogin
 * @Description 支持手机号,密码,图形验证码或者手机号,短信验证码
 * @Author khr
 * @Date 2023/5/8 9:46
 */

// @Summary 登录
// @Description 登录
// @Tags 用户信息
// @Param user body reqDto.NewUserLogin true "用户信息"
// @Success 200 {Object} UserInfoResponse "返回结果"
// @Failure 400 {Object} ErrorResponse "请求错误"
// @Router /api/user/login [post]
func newLogin(c *gin.Context) {
	//res := global.NewResult(c)
	var loginDto reqDto.NewUserLogin
	if err := c.BindJSON(&loginDto); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.Error(err)
			return
		}
		c.Error(errs)
		return
	}

	// 建立WaitGroup
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//userLoginErr := make(chan error)
	//userLoginData := make(chan resDto.TokenAndExp)
	//endErr, endData := userService.NewUserLogin(loginDto)
	//fmt.Println(endErr, endData)
	//if endErr != nil {
	//	c.Error(endErr)
	//	return
	//}
	c.Set("res", "endData")
}

/*
 * @MethodName signUser
 * @Description 手机号,密码图片验证
 * @Author khr
 * @Date 2023/5/8 14:01
 */

func signUser(c *gin.Context) {
	var sign reqDto.SignUser
	if err := c.BindJSON(&sign); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.Error(err)
			return
		}
		c.Error(errs)
		return
	}
	//err, res := userService.UserSin(sign)
	//if err != nil {
	//	c.Error(err)
	//	return
	//}
	c.Set("res", "res")
}

//func rejisterUser(c *gin.Context) {
//	res := global.NewResult(c)
//	var add reqDto.AddUser
//	if err := c.BindJSON(&add); err != nil {
//		errs, ok := err.(validator.ValidationErrors)
//		if !ok {
//			res.Error(http.StatusInternalServerError, err.Error())
//			return
//		}
//		res.Error(http.StatusBadRequest, global.Translate(errs))
//		return
//	}
//	resErr := make(chan error)
//	resData := make(chan string)
//	go userService.UserRejist(add, resErr, resData)
//	endErr := <-resErr
//	endData := <-resData
//	if endErr != nil {
//		res.Err(endData)
//		return
//	}
//	res.Success(endData)
//	return
//	//bol, msg := userService.UserRejist(add)
//	//if !bol {
//	//	res.Err(msg)
//	//	return
//	//}
//	//res.Success(msg)
//	//return
//}

func postMessage(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	fmt.Println(name, pwd, "传递参数")
	c.String(http.StatusOK, fmt.Sprintf("name:%s ,pwd:%s,type:%s", name, pwd, types))
}

// @Summary 获取用户信息
// @Description 获取指定用户的信息
// @Tags 用户信息
// @Param id path int true "用户ID"
// @Success 200 {Object} UserInfoResponse "返回结果"
// @Failure 400 {Object} ErrorResponse "请求错误"
// @Router /api/user/{id} [get]
func getUserInfo(c *gin.Context) {
	res := global.NewResult(c)
	user, _ := c.Get("user")
	fmt.Println("request", user)
	res.Success(gin.H{
		"message": "hello gin",
		"request": user,
	})
	return

}

func updateUser(c *gin.Context) {
	result := global.NewResult(c)

	result.Success(util.MODIFICATION_SUCCESSE)
	return
}
