package adminService

import (
	"gin-web/src/dto/comDto"
	"gin-web/src/dto/reqDto"
	"gin-web/src/dto/resDto"
	"gin-web/src/global"
	"gin-web/src/pojo"
	"gin-web/src/service/userService"
	"gin-web/src/util"
)

//type AdminService struct {
//}

var admin pojo.Admin
var err error
var adminInfo = &resDto.AdminInfo{}
var permissionInfo = &resDto.PermissionInfo{}

// 引入dao层
var (
	permissionServiceImpl = pojo.RbacPermission()
	roleServiceImpl       = pojo.RbacRule()
	adminServiceImpl      = pojo.AdminServiceImpl()
)

/*
*
反射控制层登录参数，查询数据库账号是否相同，
比对密码一致性，将用户信息存入jwt令牌中，签发令牌和过期时间
*/
func AdminLogin(list reqDto.AdminLogin) (err error, tokenAndExp interface{}) {
	var ad = &pojo.Admin{}
	switch list.Method {
	case "name":
		ad, err = adminServiceImpl.CheckByName(list.Name)
	case "account":
		ad, err = adminServiceImpl.CheckByAccount(list.Account)
	default:
		return err, util.METHOD_NOT_FILLED_ERROR
	}
	if err != nil {
		return err, util.ACCOUT_NOT_EXIST_ERROR
	}
	pwd, deerr := util.DePwdCode(ad.Password, ad.Salt)
	if deerr != nil {
		return err, util.PASSWORD_RESOLUTION_ERROR
	}
	if pwd == "" {
		return err, util.PASSWORD_RESOLUTION_ERROR
	}
	if pwd != list.Password {
		return err, util.AUTH_LOGIN_PASSWORD_ERROR
	}
	_, role_name := roleServiceImpl.FindRoleName(uint(ad.Role))
	existOldToken := util.ExistRedis(ad.AccessToken)
	tokenKey := util.Rand6String6()
	var token string
	var exptime string

	stringTokenData := comDto.TokenClaims{
		Id:       ad.ID,
		Name:     ad.Name,
		Role:     ad.Role,
		Account:  ad.Account,
		RoleName: role_name.Name,
	}
	switch list.Revoke {
	case true:
		if existOldToken {
			util.DelRedis(ad.AccessToken) //清除token
		}
		token, exptime = util.SignToken(stringTokenData, global.AdminLoginTime*global.DayTime)
		err = adminServiceImpl.UpdateToken(tokenKey, ad.ID)
		if err != nil {
			return err, util.AUTH_LOGIN_ERROR
		}
		redisDate := reqDto.LoginRedisDate{
			Token:   token,
			Exptime: exptime,
		}
		util.SetRedis(tokenKey, util.Marshal(redisDate), global.AdminLoginTime)
		tokenAndExp = resDto.TokenAndExp{
			token,
			exptime,
		}
	case false:
		if existOldToken {
			tokenValue := util.GetRedis(ad.AccessToken)
			mp := make(map[string]interface{})
			_, cs := util.UnMarshal([]byte(tokenValue), &mp)
			return nil, cs
		}
		//token过期时
		token, exptime = util.SignToken(stringTokenData, global.AdminLoginTime*global.DayTime)
		err = adminServiceImpl.UpdateToken(tokenKey, ad.ID)
		if err != nil {
			return err, util.AUTH_LOGIN_ERROR
		}
		redisDate := reqDto.LoginRedisDate{
			Token:   token,
			Exptime: exptime,
		}
		util.SetRedis(tokenKey, util.Marshal(redisDate), global.AdminLoginTime)
		tokenAndExp = resDto.TokenAndExp{
			token,
			exptime,
		}
		break
	}
	return nil, tokenAndExp
}

func AdminInfo(id int, name string) (bool, *resDto.AdminInfo) {
	//var adminInfo = resDto.AdminInfo{}
	//var ok bool
	adminInfo, err = adminServiceImpl.AdminInfo(id, name)
	if err != nil {
		return false, nil
	}
	return true, adminInfo
}

// 分页模糊查询管理员
func AdminList(list reqDto.AdminList, resErr chan error, resData chan interface{}) {
	reslist := &resDto.CommonList{}
	reslist, err = adminServiceImpl.AdminList(list)
	if err != nil {
		resErr <- err
	}
	resData <- reslist
	//res := adminServiceImpl.AdminList(list)
	//fmt.Println("adminlist:", "res")
	//return "res"
}

// 增加admin
func AdminAdd(add reqDto.AddAdmin, resErr chan error, resData chan interface{}) {
	add.Salt = util.RandAllString()
	var pwd = add.Password
	//校验是否有密码，没有则为123456
	if add.Password == "" {
		pwd = string(123456)
	}
	//调用加密方法
	enPwd, _ := util.EnPwdCode(pwd, add.Salt)
	//加密密码
	add.Password = enPwd
	//检查名称是否重复
	if add.Name != "" {
		_, err = adminServiceImpl.CheckByName(add.Name)
		if err != nil {
			resErr <- err
			resData <- util.NAME_EXIST_ERROR
			//return false, util.NAME_EXIST_ERROR
		}
	}
	if add.Name == "" {
		add.Name = "暂未命名"
	}
	_, err = adminServiceImpl.CheckByAccount(add.Account)
	if err != nil {
		resErr <- err
		resData <- util.ACCOUNT_EXIST_ERROR
		//return false, util.ACCOUNT_EXIST_ERROR
	}
	ad := pojo.Admin{
		Salt:     add.Salt,
		Password: add.Password,
		Name:     add.Name,
		Account:  add.Account,
		Role:     add.Role}
	err = adminServiceImpl.AddAdmin(ad)
	if err != nil {

		resErr <- err
		resData <- util.ADD_ERROR
		//return true, util.ADD_SUCCESS
	} else {
		resErr <- nil
		resData <- util.ADD_SUCCESS
		//return false, util.ADD_ERROR
	}
}

// 调用userservice服务层的服务
func UserList(list reqDto.UserList) interface{} {
	res := userService.UserList(list)
	return res
}

// 登出
func AdminLogout() {
	util.DelRedis(admin.AccessToken)
}

// 权限列表
func PermissionList(list reqDto.PermissionList) (error, interface{}) {
	var resList = &resDto.CommonList{}
	err, resList = permissionServiceImpl.FindPermissionList(list)
	if err != nil {
		return err, nil
	}
	return nil, resList
}

/*权限增加*/
func PermissionAdd(permission reqDto.PermissionAdd) error {
	per := pojo.Permission{
		Host:            permission.Host,
		Path:            permission.Path,
		AuthorizedRoles: permission.AuthorizedRoles,
		ForbiddenRoles:  permission.ForbiddenRoles,
		Method:          permission.Method,
		AllowAnyone:     permission.AllowAnyone,
	}
	return permissionServiceImpl.AddPermission(per)
}

/*权限修改*/
//func PermissionUpdate(permission reqDto.PermissionUpdate) error {
//	return permissionServiceImpl.SavePermission(permission)
//
//}

// 权限删除
func Permissiondel(id int) (bool, string) {
	err, _ = permissionServiceImpl.FindPermissionById(id)
	if err == nil {
		if err = permissionServiceImpl.DelPermission(id); err != nil {
			return true, util.DELETE_SUCCESS
		}
		return false, util.PERMISSION_NOT_FOUND_ERROR
	}
	return false, util.PERMISSION_NOT_FOUND_ERROR
}

/*权限详情*/
func PermissionIndo(id int) (bool, interface{}) {
	err, permissionInfo = permissionServiceImpl.FindPermissionById(id)
	if err != nil {
		return false, err
	}
	return true, permissionInfo
}
