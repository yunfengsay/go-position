package models

import (
	"fmt"
	"position/models/class"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=zhangyunfeng password='' host=127.0.0.1 port=5432 dbname=beefer sslmode=disable", 30)
	orm.RegisterModel(new(class.User))
	orm.RunSyncdb("default", false, true)
	fmt.Println("数据库执行")
}
