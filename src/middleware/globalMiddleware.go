package middleWare

import (
	"HelloGin/src/dto/comDto"
	"HelloGin/src/global"
	"HelloGin/src/util"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"sync"
	"time"
)

var ii comDto.TokenClaims

//var permissionServiceImpl = pojo.RbacPermission()

// 定义一个记录请求次数的map
var requestCounts = make(map[string]int)

// 全局路由中间件检测token
func GolbalMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("token认证开始执行")
		//t := time.Now()
		requestUrl := c.Request.URL.String()
		//路径模糊匹配
		if !util.FuzzyMatch(requestUrl, global.ReuqestPaths) {
			//请求头是否携带token
			judge := util.AnalysyToken(c)
			if !judge {
				c.AbortWithStatusJSON(http.StatusUnauthorized, util.NO_AUTHORIZATION)
				return
			}
			ii = util.ParseToken(c.GetHeader("Authorization"))
			c.Set("role_name", ii.RoleName)
		}
		c.Next()
		//ts := time.Since(t)
		//fmt.Println("time", ts)
		//fmt.Println("token认证执行结束")

	}
}

// 权限路由中间件
//func AuthMiddleWare() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		requestUrl := c.Request.URL.String()
//		reqUrl := strings.Split(requestUrl, "/api/")
//		rolename := global.RoleName
//		paths := global.ReuqestPaths
//		pathIsExist := util.ExistIn(reqUrl[1], paths)
//		//登录跳过权限验证
//		if !pathIsExist {
//			//验证身份
//			_, y := c.Get("ok")
//			//通过身份验证
//			if !y {
//				c.AbortWithStatusJSON(http.StatusUnauthorized, util.NO_AUTH_ERROR)
//				return
//			} else {
//				roleName := c.GetString("role_name")
//				role := c.GetInt("role")
//				if !util.ExistIn(roleName, rolename) {
//					err, permission := permissionServiceImpl.FindPermissionByPath(reqUrl[1])
//					if err != nil {
//						c.AbortWithStatusJSON(http.StatusAccepted, util.INSUFFICIENT_PERMISSION_ERROR)
//						return
//					}
//					allowRole := permission.AuthorizedRoles
//					roleList := strings.Split(allowRole, ",")
//					roleExist := util.ExistIn(string(role), roleList)
//					if !roleExist {
//						//c.Abort()
//						//fmt.Println("请求地址不包含该权限权限")
//						c.AbortWithStatusJSON(http.StatusAccepted, util.INSUFFICENT_PERMISSION)
//						//res.Err(util.INSUFFICENT_PERMISSION)
//						return
//					}
//				}
//				fmt.Println("检测到是超级管理员，可以直接操作，不需要判断")
//			}
//		}
//	}
//}

// 全局日志中间件
func LoggerMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		requestBody, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		rbody := string(requestBody)
		query := c.Request.URL.RawQuery
		c.Next() // 调用该请求的剩余处理程序
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000))))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		url := c.Request.RequestURI
		Log := global.Logger.WithFields(
			logrus.Fields{
				"SpendTime": spendTime,       //接口花费时间
				"path":      url,             //请求路径
				"Method":    method,          //请求方法
				"status":    statusCode,      //接口返回状态
				"proto":     c.Request.Proto, //http请求版本
				"Ip":        clientIP,        //IP地址
				"body":      rbody,           //请求体
				"query":     query,           //请求query
				"message":   c.Errors,        //返回错误信息
			})
		if len(c.Errors) > 0 { // 矿建内部错误
			Log.Error(c.Errors.ByType(gin.ErrorTypePrivate))
		}
		if statusCode > 200 {
			Log.Error()
		} else {
			Log.Info()
		}
	}
}

var visitorMap = make(map[string]*rate.Limiter) // 存储IP地址和速率限制器的映射
var mu sync.Mutex                               // 互斥锁，保证并发安全
// IP请求次数拦截中间件
// 定义一个中间件函数，用于拦截请求并进行计数
func IPInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		if ip == "" {
			ip = c.Request.RemoteAddr
		}
		if util.ExistIn(ip, global.IpAccess) {
			c.Next()
		}
		path := c.Request.URL.Path
		//fmt.Println(ip, path)

		// 组合出 key
		key := fmt.Sprintf("request:%s:%s", ip, path)
		//fmt.Print("key", key)
		// 将请求次数 +1，并设置过期时间
		_, err := global.Redis.Incr(c, key).Result()

		if err != nil {
			// 记录日志
			fmt.Println("incr error:", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		global.Redis.Expire(c, key, time.Hour)

		// 获取当前IP在 path 上的请求次数
		accessTime, err := global.Redis.Get(c, key).Int()

		if err != nil {
			// 记录日志
			fmt.Println("get error:", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		//ip一小时内访问路径超过次数限制，拒绝访问
		if accessTime > 60 {
			requestLimit := fmt.Sprintf("request:%s:%s", ip, path)
			global.Redis.RPush(c, global.RedisReqLimitUrl, requestLimit)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			return
		}
		mu.Lock()
		_, ok := visitorMap[ip]
		var limiter = rate.NewLimiter(1, 10) // 设置限制为1个请求/秒，最多允许10个并发请求
		// 如果该IP地址不存在，则创建一个速率限制器
		if !ok {
			visitorMap[ip] = limiter
		}
		mu.Unlock()
		// 尝试获取令牌，如果没有可用的令牌则阻塞
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			return
		}
		c.Next()
	}
}

// casbin中间件
func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("权限认证开始执行")
		requestUrl := c.Request.URL.String()
		if !util.FuzzyMatch(requestUrl, global.ReuqestPaths) {
			reqUrl := strings.Split(requestUrl, "/api/")
			paths := global.ReuqestPaths
			pathIsExist := util.ExistIn(reqUrl[1], paths)
			//fmt.Println("在开放权限")
			if !pathIsExist {
				fmt.Println("不在开放权限")
				sub := "superadmin"
				dom := ""
				//sub, _ := c.Get("role_name")
				obj := c.Request.URL.Path
				act := c.Request.Method
				ok, _ := global.CasbinDb.Enforce(sub, dom, obj, act)
				fmt.Println("见电工出错", ok)
				if !ok {
					c.AbortWithStatusJSON(403, gin.H{"error": "not authorized"})
					return
				}

				fmt.Println("可以访问")
				c.Next()
			}
		}

		c.Next()
		//fmt.Println("权限认证执行结束")
	}
	//return func(next http.Handler) http.Handler {
	//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		requestUrl := r.URL.String()
	//		reqUrl := strings.Split(requestUrl, "/api/")
	//		paths := global.ReuqestPaths
	//		pathIsExist := util.ExistIn(reqUrl[1], paths)
	//		if !pathIsExist {
	//			// 获取请求的用户和路径信息
	//			subject := r.Header.Get("X-User")
	//			path := r.URL.Path
	//			method := r.Method
	//
	//			// 检查当前用户是否有访问该路径的权限
	//			if ok, _ := e.Enforce(subject, path, method); ok {
	//				// 如果有权限，执行下一个处理器
	//				next.ServeHTTP(w, r)
	//			} else {
	//				// 如果没有权限，返回错误
	//				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	//			}
	//		}
	//	})
	//}
}

// 统一返回格式
func FormatResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 判断是否有错误信息
		if len(c.Errors) > 0 {
			fmt.Println("出现错误", c.Errors.Last().Error())
			// 返回错误信息
			c.JSON(http.StatusOK, gin.H{

				"code": http.StatusInternalServerError,
				"msg":  c.Errors.Last().Error(),
				"data": "",
			})
			return
		}

		// 判断是否有返回数据
		if c.Writer.Status() == http.StatusOK {
			// 获取返回数据
			data, ok := c.Get("res")
			if ok {
				// 格式化返回数据
				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusOK,
					"msg":  "",
					"data": data,
				})
				return
			}
		}

		// 返回空数据
		c.JSON(http.StatusOK, gin.H{})
	}
}
