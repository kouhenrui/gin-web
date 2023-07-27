package admin

import (
	"fmt"
	"gin-web/src/dto/reqDto"
	"gin-web/src/global"
	"gin-web/src/service/adminService"
	"gin-web/src/service/rbacService"
	"gin-web/src/service/userService"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

func Routers(e *gin.Engine) {
	adminGroup := e.Group("/api/admin")
	{
		adminGroup.POST("/register", registerAdmin)
		adminGroup.POST("/login", adminLogin)
		adminGroup.GET("/info", getAdminInfo)
		adminGroup.GET("/logout", logout)
		adminGroup.POST("/list", adminList)
		adminGroup.POST("/users/list", userList)
		adminGroup.POST("/permission/list", permissionList)
		adminGroup.POST("/permission/add", permissionAdd)
		adminGroup.PUT("/permission/update", permissionUpdate)
		adminGroup.DELETE("/permission/del", permissionDelete) //permission/del?id=8
		adminGroup.GET("/permission/info", permissionInfo)     //permission/info?id=
		//adminGroup.POST("/group/list", groupList)
		//adminGroup.POST("/permission/add", permissionAdd)
		//adminGroup.PUT("/permission/update", permissionUpdate)
		//adminGroup.DELETE("/permission/del", permissionDelete) //permission/del?id=8
		//adminGroup.GET("/permission/info", permissionInfo)     //permission/info?id=
	}
}

// 登录接口
func adminLogin(c *gin.Context) {
	//res := global.NewResult(c)
	var lg reqDto.AdminLogin
	if err := c.BindJSON(&lg); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.Error(err)
			return
		}
		c.Error(errs)
		return
	}
	logerr, result := adminService.AdminLogin(lg)
	if logerr != nil {
		c.Error(logerr)
		return
	}
	c.Set("res", result)
	return
}

// 登出
func logout(c *gin.Context) {
	res := global.NewResult(c)
	go adminService.AdminLogout()
	res.Succ()
}

// 获取详情接口
func getAdminInfo(c *gin.Context) {
	id := c.GetInt("id")
	name := c.GetString("name")
	_, info := adminService.AdminInfo(id, name)

	c.Set("res", info)
}

// 增加管理员接口
func registerAdmin(c *gin.Context) {
	var re reqDto.AddAdmin
	if err := c.BindJSON(&re); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.Error(err)
			return
		}
		c.Error(errs)
		return
	}
	//异步线程操作
	resErr := make(chan error)
	resData := make(chan interface{})
	go adminService.AdminAdd(re, resErr, resData)
	endErr := <-resErr
	endData := <-resData
	if endErr != nil {
		c.Error(endErr)
		return
	}
	c.Set("res", endData)
	//judje, result := adminService.AdminAdd(re)
	//if judje {
	//	res.Success(result)
	//	return
	//} else {
	//	res.Err(result)
	//	return
	//}
}

// 管理员列表
func adminList(c *gin.Context) {
	var ls reqDto.AdminList
	fmt.Println("请求参数：", ls)
	if err := c.BindJSON(&ls); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.Error(err)
			return
		}
		c.Error(errs)
		return
	}
	//异步线程操作
	resErr := make(chan error)
	resData := make(chan interface{})
	go adminService.AdminList(ls, resErr, resData)
	endErr := <-resErr
	endData := <-resData
	if endErr != nil {
		c.Error(endErr)
		return
	}
	c.Set("res", endData)
	//list := adminService.AdminList(ls)
	//fmt.Println("list:", list)
	//res.Success(list)
	//return

}

// 用户列表
func userList(c *gin.Context) {
	res := global.NewResult(c)
	var ls reqDto.UserList
	if err := c.BindJSON(&ls); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.Error(http.StatusBadRequest, err.Error())
			return
		}
		res.Error(http.StatusBadRequest, global.Translate(errs))
		return
	}
	list := userService.UserList(ls)
	//list := adminService.UserList(ls)
	res.Success(list)
	return
}

// 权限列表
func permissionList(c *gin.Context) {
	var ls reqDto.PermissionList
	//fmt.Println("请求体：", ls.Take, ls.Skip)
	if err := c.BindJSON(&ls); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.Error(err)
			return
		}
		c.Error(errs)
		return
	}
	listerr, list := adminService.PermissionList(ls)
	if listerr != nil {
		c.Error(listerr)
		return
	}
	c.Set("res", list)
}

// 权限增加
func permissionAdd(c *gin.Context) {
	res := global.NewResult(c)
	var ls reqDto.PermissionAdd
	if err := c.BindJSON(&ls); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.Error(http.StatusBadRequest, err.Error())
			//c.Abort()
			return
		}
		res.Error(http.StatusBadRequest, global.Translate(errs))
		//c.Abort()
		return
	}
	err := adminService.PermissionAdd(ls)
	if err != nil {
		res.Err(err)
		return
	}
	res.Succ()
	return
}

/* 权限修改*/
func permissionUpdate(c *gin.Context) {
	res := global.NewResult(c)
	var ls reqDto.PermissionUpdate
	if err := c.BindJSON(&ls); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.Error(http.StatusBadRequest, err.Error())
			return
		}
		res.Error(http.StatusBadRequest, global.Translate(errs))
		return
	}

	err := rbacService.UpdatePermission(ls)
	if err != nil {
		res.Err(err)
		return
	}
	res.Succ()
	return
}

// 权限删除
func permissionDelete(c *gin.Context) {
	res := global.NewResult(c)
	id := c.Query("id")
	i, _ := strconv.Atoi(id)
	ok, result := adminService.Permissiondel(i)
	if ok {
		res.Success(result)
		return
	}
	res.Err(result)
	return
}

/*权限详情*/
func permissionInfo(c *gin.Context) {
	res := global.NewResult(c)
	id := c.Query("id")
	i, _ := strconv.Atoi(id) //string转int
	ok, result := adminService.PermissionIndo(i)
	if ok {
		res.Success(result)
		return
	}
	res.Err(result)
	return
}

/*
 * @MethodName roleList
 * @Description 角色列表
 * @Author Acer
 * @Date 2023/4/3 15:52
 */

func roleList(c *gin.Context) {
	res := global.NewResult(c)
	var rulereq reqDto.RuleList
	if err := c.BindJSON(&rulereq); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.Error(http.StatusBadRequest, err.Error())
			//c.Abort()
			return
		}
		res.Error(http.StatusBadRequest, global.Translate(errs))
		//c.Abort()
		return
	}

}

/*角色增加*/

/*组列表*/

/*组增加*/
