package main

import (
	_ "position/routers"

	_ "position/models"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

// func init() {
// 	orm.RegisterDriver("postgres", orm.DRPostgres)
// 	orm.RegisterDataBase("default", "postgres", "user=zhangyunfeng password='' host=127.0.0.1 port=5432 dbname=beefer sslmode=disable", 30)
// 	orm.RunSyncdb("default", false, true)
// }

func main() {
	// orm.Debug = true
	// o := orm.NewOrm()
	// o.Using("default")
	// u := new(models.User)
	// u.Username = "zhangyunfeng"
	// u.Age = 25
	// u.Summary = "i am running"
	// u.Email = "fdoctor00@gmail.com"
	// u.Auth = 1
	// id, err := o.Insert(u)
	// fmt.Printf("ID: %d, ERR: %v\n", id, err)
	beego.SetLogger("file", `{"filename": "./logs/position.run.log"}`)
	beego.BeeLogger.DelLogger("console")

	beego.SetLogFuncCall(true)

	beego.Run()
}
