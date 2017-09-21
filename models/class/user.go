package class

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// 完成User类型定义
type User struct {
	Id        int `orm:"pk"` // 设置为主键，字段Id, Password首字母必须大写
	User_name string
	Age       int
	Password  string
	Email     string
	Auth      int
	Summary   string
	Open_id string
	Nick_name string
	Province string
	Language string
	Avatar_url string
	Create_date string
	Update_date string
	Gender	int
	Black_hole string
	Phone	string
}

func (u *User) ReadDB() (err error) {
	o := orm.NewOrm()
	err = o.Read(u)
	return err
}

func (u *User) Create() (err error) {
	o := orm.NewOrm()
	fmt.Println("Create success!")
	_, _ = o.Insert(u)
	return err
}

func (u *User) Update() (err error) {
	o := orm.NewOrm()
	_, err = o.Update(u)
	return err
}
