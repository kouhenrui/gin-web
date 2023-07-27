package userService

import "gin-web/src/dto/reqDto"

// import (
//
//	"HelloGin/src/dto/comDto"
//	"HelloGin/src/dto/resDto"
//	"HelloGin/src/global"
//	"HelloGin/src/service/casbinService"
//	"HelloGin/src/util"
//	"errors"
//	"fmt"
//	"github.com/mitchellh/mapstructure"
//	"login/src/dto/reqDto"
//	"login/src/pojo"
//
// )
//
// var user pojo.User
// var userServiceImpl = pojo.UserServiceImpl()
// var userInfo = &resDto.UserInfo{}
// var err error
//
// /*
// * @MethodName NewUserLogin
// * @Description 支持图片验证
// * @Author khr
// * @Date 2023/5/8 14:06
// */
//
//	func Login(dto reqDto.Login) (error, interface{}) {
//		dtoErr, dtoUser := userServiceImpl.CheckPhone(dto.Phone)
//		if dtoErr != nil {
//			return dtoErr, nil
//		}
//		enpwd := dtoUser.Password
//		salt := dtoUser.Salt
//		pwd, pwdErr := util.DePwdCode(enpwd, salt)
//		if pwdErr != nil {
//			return errors.New(util.PASSWORD_RESOLUTION_ERROR), nil
//		}
//		if pwd != dto.Password {
//			return errors.New(util.AUTH_LOGIN_PASSWORD_ERROR), nil
//		}
//		switch dto.Method {
//		case "captcha":
//			capres := util.VerifyCaptcha(dto.Captcha)
//			if !capres {
//				return errors.New(util.VERIFY_CODE_ERROR), nil
//			}
//			break
//		case "message":
//			break
//		default:
//			return errors.New(util.METHOD_NOT_FILLED_ERROR), nil
//			break
//		}
//		//获取redis缓存
//		existOldToken := util.ExistRedis(dtoUser.AccessToken)
//		tokenKey := util.Rand6String6()
//		var token string
//		var exptime string
//		ruleErr, roleName := casbinService.CheckRuleName(uint(dtoUser.Role))
//		if ruleErr != nil {
//			roleName = ""
//		}
//		stringTokenData := comDto.TokenClaims{
//			Id:       dtoUser.ID,
//			Name:     dtoUser.Name,
//			Role:     dtoUser.Role,
//			Phone:    dtoUser.Phone,
//			Account:  dtoUser.Account,
//			RoleName: roleName,
//		}
//		tokenAndExp := resDto.TokenAndExp{}
//		switch dto.Revoke {
//		case true:
//			if existOldToken {
//				util.DelRedis(dtoUser.AccessToken) //清除token
//			}
//			token, exptime = util.SignToken(stringTokenData, global.UserLoginTime*global.DayTime)
//			err = userServiceImpl.UpdateToken(tokenKey, dtoUser.ID)
//			if err != nil {
//				return err, nil
//			}
//			redisDate := reqDto.LoginRedisDate{
//				Token:   token,
//				Exptime: exptime,
//			}
//			util.SetRedis(tokenKey, util.Marshal(redisDate), global.UserLoginTime)
//			tokenAndExp.Token = token
//			tokenAndExp.Exptime = exptime
//			return nil, tokenAndExp
//			break
//		case false:
//			if existOldToken {
//				tokenValue := util.GetRedis(dtoUser.AccessToken)
//				mp := make(map[string]interface{})
//				_, cs := util.UnMarshal([]byte(tokenValue), &mp)
//				er := mapstructure.Decode(cs, &tokenAndExp)
//				if er != nil {
//					return err, nil
//				}
//				break
//			} else {
//				token, exptime = util.SignToken(stringTokenData, global.UserLoginTime*global.DayTime)
//				err = userServiceImpl.UpdateToken(tokenKey, dtoUser.ID)
//				if err != nil {
//
//					return errors.New(util.AUTH_LOGIN_ERROR), nil
//				}
//				redisDate := reqDto.LoginRedisDate{
//					Token:   token,
//					Exptime: exptime,
//				}
//				_ = util.SetRedis(tokenKey, util.Marshal(redisDate), global.UserLoginTime)
//				tokenAndExp.Token = token
//				tokenAndExp.Exptime = exptime
//				return nil, tokenAndExp
//				break
//			}
//		default:
//			return errors.New(util.REQUEST_NOT_EXIST), nil
//			break
//		}
//		return nil, tokenAndExp
//
// }
//
//	func UserSin(dto reqDto.SignUser) (error, interface{}) {
//		capres := util.VerifyCaptcha(dto.Captcha)
//		if !capres {
//			return errors.New(util.VERIFY_CODE_ERROR), nil
//		}
//		err, _ = userServiceImpl.CheckPhone(dto.Phone)
//		if err != nil {
//			return err, nil
//		}
//		//校验是否有密码，没有则为123456
//		var pwd string
//		if dto.Password == "" {
//			pwd = string(123456)
//		}
//		salt := util.RandAllString()
//		//调用加密方法
//		password, _ := util.EnPwdCode(pwd, salt)
//		user.Phone = dto.Phone
//		user.Password = password
//		user.Salt = salt
//		err = userServiceImpl.AddUser(user)
//		fmt.Println(err, "写入错误")
//		if err != nil {
//			return err, nil
//		}
//		return nil, util.ADD_SUCCESS
//	}
//
// // 用户列表
func UserList(list reqDto.UserList) interface{} {
	//res := userServiceImpl.UserList(list)
	//return res
	return "success"
}

//
//// 登录
//func UserLogin(list reqDto.UserLogin, resErr chan error, resData chan resDto.TokenAndExp) {
//	//fmt.Println("查询的数据", list.Method)
//	var userlogin = &pojo.User{}
//	switch list.Method {
//	case "name":
//		err, userlogin = userServiceImpl.CheckByName(list.Name)
//	case "account":
//		err, userlogin = userServiceImpl.CheckByAccount(list.Account)
//	default:
//		resErr <- errors.New(util.METHOD_NOT_FILLED_ERROR)
//		//return false, util.METHOD_NOT_FILLED_ERROR
//	}
//	if err != nil {
//		resErr <- errors.New(util.ACCOUT_NOT_EXIST_ERROR)
//		//return false, util.ACCOUT_NOT_EXIST_ERROR
//	}
//	enpwd := userlogin.Password
//	salt := userlogin.Salt
//	pwd, deerr := util.DePwdCode(enpwd, salt)
//	if deerr != nil {
//		resErr <- errors.New(util.PASSWORD_RESOLUTION_ERROR)
//		//return false, util.PASSWORD_RESOLUTION_ERROR
//	}
//	if pwd != list.Password {
//		resErr <- errors.New(util.AUTH_LOGIN_PASSWORD_ERROR)
//		//return false, util.AUTH_LOGIN_PASSWORD_ERROR
//	}
//	existOldToken := util.ExistRedis(userlogin.AccessToken)
//	tokenKey := util.Rand6String6()
//	var token string
//	var exptime string
//	stringTokenData := comDto.TokenClaims{
//		Id:      userlogin.ID,
//		Name:    userlogin.Name,
//		Role:    userlogin.Role,
//		Phone:   userlogin.Phone,
//		Account: userlogin.Account,
//		//RoleName: roleName,
//	}
//	tokenAndExp := resDto.TokenAndExp{}
//	switch list.Revoke {
//	case true:
//		if existOldToken {
//			util.DelRedis(userlogin.AccessToken) //清除token
//		}
//		token, exptime = util.SignToken(stringTokenData, global.UserLoginTime*global.DayTime)
//		err = userServiceImpl.UpdateToken(tokenKey, userlogin.ID)
//		if err != nil {
//			resErr <- errors.New(util.AUTH_LOGIN_ERROR)
//			//return false, util.AUTH_LOGIN_ERROR
//		}
//		redisDate := reqDto.LoginRedisDate{
//			Token:   token,
//			Exptime: exptime,
//		}
//		util.SetRedis(tokenKey, util.Marshal(redisDate), global.UserLoginTime)
//		tokenAndExp.Token = token
//		tokenAndExp.Exptime = exptime
//		break
//	case false:
//		if existOldToken {
//			tokenValue := util.GetRedis(userlogin.AccessToken)
//			mp := make(map[string]interface{})
//			_, cs := util.UnMarshal([]byte(tokenValue), &mp)
//			er := mapstructure.Decode(cs, &tokenAndExp)
//			if er != nil {
//				fmt.Println("redis缓存转化错误", er)
//				resErr <- er
//				//resErr <- nil
//				//resData <- tokenAndExp
//			}
//			break
//		} else {
//			token, exptime = util.SignToken(stringTokenData, global.UserLoginTime*global.DayTime)
//			err = userServiceImpl.UpdateToken(tokenKey, userlogin.ID)
//			//fmt.Println(err, "更新token错误")
//			if err != nil {
//				resErr <- errors.New(util.AUTH_LOGIN_ERROR)
//				//return false, util.AUTH_LOGIN_ERROR
//			}
//			redisDate := reqDto.LoginRedisDate{
//				Token:   token,
//				Exptime: exptime,
//			}
//			util.SetRedis(tokenKey, util.Marshal(redisDate), global.UserLoginTime)
//			tokenAndExp.Token = token
//			tokenAndExp.Exptime = exptime
//			break
//		}
//
//	}
//	resErr <- nil
//	resData <- tokenAndExp
//}
//
//// 注册
//func UserRejist(list reqDto.AddUser, resErr chan error, resData chan string) {
//	list.Salt = util.RandAllString()
//	var pwd = list.Password
//	//校验是否有密码，没有则为123456
//	if list.Password == "" {
//		pwd = string(123456)
//	}
//	//调用加密方法
//	enPwd, _ := util.EnPwdCode(pwd, list.Salt)
//	//加密密码
//	list.Password = enPwd
//	//return true, nil
//	//检查名称是否重复
//	if list.Name != "" {
//		err, _ := userServiceImpl.CheckByName(list.Name)
//		if err != nil {
//			resErr <- err
//			resData <- util.NAME_EXIST_ERROR
//			//return false, util.NAME_EXIST_ERROR
//		}
//	}
//	if list.Name == "" {
//		list.Name = "暂未命名"
//	}
//	erro, _ := userServiceImpl.CheckByAccount(list.Account)
//	fmt.Println(erro, "账号存在的错误")
//	if erro != nil {
//		resErr <- erro
//		resData <- util.ACCOUNT_EXIST_ERROR
//		//return false, util.ACCOUNT_EXIST_ERROR
//	}
//	ad := pojo.User{
//		Salt:     list.Salt,
//		Password: list.Password,
//		Name:     list.Name,
//		Account:  list.Account,
//		Phone:    list.Phone,
//		Role:     list.Role}
//	er := userServiceImpl.AddUser(ad)
//	if er != nil {
//		resErr <- er
//		resData <- util.ADD_ERROR
//		//return true, util.ADD_SUCCESS
//	} else {
//		resErr <- nil
//		resData <- util.ADD_SUCCESS
//		//return false, util.ADD_ERROR
//	}
//}
//
////func UserByNameAndAccount(query string) bool {
////	result := db.Where(query).Take(&u)
////	if result.Error != nil {
////		return false
////	}
////	return true
////}
////
//////func JudgeUserExist(name string, account string) (a chan bool) {
//////	//ua := pojo.User{}
//////	db.Select("id", "name", "account").Where("name = ? or account=?", name, account).First(&u)
//////	log.Println(u, "打印数据公告", u.ID, "查询id")
//////	if u.ID == 0 {
//////		a <- true
//////		return
//////	}
//////	a <- false
//////	return
//////}
//////
