package global

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	_ "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"log"
	"strings"
)

/**
 * @ClassName casbin
 * @Description TODO
 * @Author khr
 * @Date 2023/4/24 14:25
 * @Version 1.0
 */

/*
 * @MethodName KeyMatchFunc
 * @Description 正则匹配
 * @Author khr
 * @Date 2023/4/24 14:26
 */
func check(sub, obj, act string) {
	ok, _ := CasbinDb.Enforce(sub, obj, act)

	//fmt.Println(er, "err")
	if ok {
		fmt.Printf("%s CAN %s %s in %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s in %s\n", sub, act, obj)
	}
}
func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[2].(string)
	return KeyMatch(name1, name2), nil
}
func KeyMatch(key1, key2 string) bool {
	i := strings.Index(key2, "*")
	if i == -1 {
		return key1 == key2
	}

	if len(key1) > i {
		return key1[:i] == key2[:i]
	}

	return key1 == key2[:i]
}

var CasbinDb *casbin.Enforcer

func CasbinInit() {
	var err error
	//fmt.Printf("打印权限", casbinCon)
	db := CasbinConfig.UserName + ":" + CasbinConfig.PassWord + "@tcp(" + CasbinConfig.HOST + ":" + CasbinConfig.Port + ")/"
	//db加库名可以指定使用表或者自动创建表
	//a, aerr := gormadapter.NewAdapter(CasbinConfig.Type, db,true)//自己创建表
	a, aerr := gormadapter.NewAdapter(CasbinConfig.Type, db)
	CasbinDb, err = casbin.NewEnforcer("auth_model.conf", a)
	//fmt.Println("e", e)
	if aerr != nil {
		fmt.Printf("权限表为创建，错误原因：%s", aerr)
	}
	if err != nil {
		fmt.Println("加载模型出现错误", err)
	}
	log.Printf("权限初始化成功")
	//使用模糊匹配路径
	CasbinDb.AddFunction("regexMatch", RegexMatchFunc)
	//创建表

	//e.AddFunction("my_func", KeyMatchFunc)
	//check(CasbinDb, "dajun", "root", "data1", "all")
	//check(e, "lili", "dev", "data2", "read")
	//check(e, "dajun", "tenant1", "data1", "read")
	//check(e, "dajun", "tenant2", "data2", "read")
	check("superadmin", "", "")
}

// 正则匹配函数
func RegexMatchFunc(args ...interface{}) (interface{}, error) {
	return util.RegexMatch(args[0].(string), args[1].(string)), nil
}
