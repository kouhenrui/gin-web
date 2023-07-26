package routers

import (
	middleWare "gin-web/src/middleware"
	"gin-web/src/util"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/unrolled/secure"
	"log"
	"net/http"
)

type Option func(engine *gin.Engine)

var options = []Option{}

func Include(opts ...Option) {
	options = append(options, opts...)
}
func InitRoute() *gin.Engine {
	log.Println("路由初始化调用")
	gin.ForceConsoleColor()
	r := gin.New()
	r.Use(Cors())                               //民间跨域
	r.StaticFS("/img", http.Dir("./static"))    //加载静态资源，一般是上传的资源，例如用户上传的图片
	r.StaticFS("/dynamic", http.Dir("./video")) //加载静态资源，一般是上传的资源，例如用户上传的图片
	//r.Use(TlsHandler())                         //转换为https协议
	r.Use(middleWare.GolbalMiddleWare()) //全局中间件
	r.Use(middleWare.LoggerMiddleWare()) //日志中间件
	r.Use(middleWare.IPInterceptor())    //请求ip次数限制
	r.Use(middleWare.CasbinMiddleware()) //使用casbin鉴权中间件
	// 初始化 swag
	//docs.SwaggerInfo.Title = "gin swag"
	//docs.SwaggerInfo.Description = "Your API description"
	//docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "localhost:8888"
	//docs.SwaggerInfo.BasePath = "/"
	//docs.SwaggerInfo.Schemes = []string{"https"}

	r.GET("/api/swag/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 注册swagger中间件
	//r.Use(middleWare.AuthMiddleWare())   //身份认证
	//r.Use(middleWare.CasbinMiddleware(global.CasbinDb)) //casbin中间件
	r.Use(middleWare.FormatResponse()) //统一返回格式
	r.Use(middleWare.Recover)          //错误捕捉
	r.Use(gin.Recovery())              //系统提供的错误捕捉
	r.NoRoute(HandleNotFound)          //路由未找到
	r.NoMethod(HandleNotAllowed)       //方法未找到
	r.MaxMultipartMemory = 64 << 20    //64Mb
	for _, ii := range options {
		ii(r)
	} //挂载模块
	return r

}

// 404
func HandleNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code": http.StatusNotFound, "msg": util.RESOURCE_NOT_FOUND_ERROR,
	})
	return
}

func HandleNotAllowed(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{
		"code": http.StatusMethodNotAllowed, "msg": util.REQUEST_METHOD_NOT_ALLOWED_ERROE,
	})
	return
}

func Cors() gin.HandlerFunc {
	//
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8080",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
