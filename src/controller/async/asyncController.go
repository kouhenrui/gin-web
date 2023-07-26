package async

import (
	"HelloGin/src/global"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Routers(e *gin.Engine) {

	asyncGroup := e.Group("/api/ync")

	async := NewAsyncController()
	asyncGroup.POST("/async", async.asyncTest)
	asyncGroup.POST("/sync", async.syncTest)
	asyncGroup.GET("/test", async.test)

}

type AsyncController struct{}

func NewAsyncController() AsyncController {
	return AsyncController{}
}

func (g *AsyncController) asyncTest(c *gin.Context) {
	result := global.NewResult(c)
	//var slice = []int{1, 2, 3, 4, 5}
	//slice[6] = 6
	copyContext := c.Copy()
	go func() {
		time.Sleep(3 * time.Second)
		log.Println("异步执行" + copyContext.Request.URL.Path)

	}()
	result.Success("success")
	return
	//c.JSON(200, gin.H{"message": "success"})
}
func (g *AsyncController) syncTest(c *gin.Context) {
	time.Sleep(3 * time.Second)
	log.Println("同步执行" + c.Request.URL.Path)
	c.JSON(200, gin.H{"message": "success"})
}
func (g *AsyncController) test(c *gin.Context) {
	time.Sleep(3 * time.Second)
	log.Println("同步执行" + c.Request.URL.Path)
	c.JSON(200, gin.H{"message": "success"})
}
