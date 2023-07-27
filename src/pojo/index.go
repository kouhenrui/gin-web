package pojo

import (
	"gin-web/src/dto/resDto"
	"gin-web/src/global"
	"log"
)

// 数据库生成表
var db = global.Db
var reslist = resDto.CommonList{}
var count int64
var (
	userpojo  = &User{}
	adminpojo = &Admin{}
	rbac_rule = &Rule{}
	rbac_per  = &Permission{}
	group     = &Group{}
)

func init() {
	db.AutoMigrate(
		user,
		adminpojo,
		rbac_rule,
		rbac_per,
		group,
	)
	log.Printf("表结构同步成功")
}

//type Base struct {
//	ID        uint           `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`
//	CreatedAt time.Time      `json:"created_at"`
//	UpdatedAt time.Time      `json:"updated_at"`
//	DeletedAt gorm.DeletedAt `gorm:"index"`
//}
