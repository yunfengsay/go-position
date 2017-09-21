package class
import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Location struct {
	Id string `orm:"pk"`
	L_type string
	Black_hole string
	Content	string
	Position [2]string
}

func (u *Location) ReadDB() (err error) {
	o := orm.NewOrm()
	err = o.Read(u)
	return err
}

func (u *Location) Create() (err error) {
	o := orm.NewOrm()
	fmt.Println("Create success!")
	_, _ = o.Insert(u)
	return err
}

func (u *Location) Update() (err error) {
	o := orm.NewOrm()
	_, err = o.Update(u)
	return err
}
