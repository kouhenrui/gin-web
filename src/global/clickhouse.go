package global

/**
 * @ClassName clickhouse
 * @Description TODO
 * @Author khr
 * @Date 2023/5/6 9:58
 * @Version 1.0
 */
import (
	"database/sql"
	_ "github.com/ClickHouse/clickhouse-go"
	"log"
	"sync"
)

var ClickhouseDb *sql.DB
var dbOnce sync.Once

func ClickhouseInit() {
	var err error
	db := "tcp://" + CasbinConfig.HOST + ":" + CasbinConfig.Port + "?" + "username=default&password=" + CasbinConfig.PassWord
	// 连接ClickHouse数据库
	dbOnce.Do(func() {
		ClickhouseDb, err = sql.Open("clickhouse", db)
		if err != nil {
			return
		}
		ClickhouseDb.SetMaxIdleConns(10)
		ClickhouseDb.SetMaxOpenConns(100)
		ClickhouseDb.SetConnMaxLifetime(30)
	})
	defer ClickhouseDb.Close()
	log.Printf("click初始化连接成功")
}
