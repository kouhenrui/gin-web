package global

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

/**
 * @ClassName cron
 * @Description 定时任务
 * @Author khr
 * @Date 2023/4/14 9:19
 * @Version 1.0
 */
var CronTesk *cron.Cron

func init() {
	//CronTesk = cron.New()

	//CronTesk.AddFunc("0 * * * * *", addCron1)
	////CronTesk.Start()
	//log.Println("定时任务初始化成功")
}
func addCron1() {
	//util.DtoToStruct(reqDto.RuleList{}, pojo.Rule{})
	fmt.Println("Task executed at", time.Now())

}
