package controllers

import (
	"fmt"
	"position/models/class"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	beego.Controller
}
type UserJson struct {
	code    int
	message string
	data    map[string]interface{}
}
type DataJson struct {
}

func (c *UserController) PageLogin() {
	c.TplName = "login.html"
}

func (c *UserController) Register() {
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println("this is id and password")
	fmt.Println(username, password)

	valid := validation.Validation{}
	valid.Required(username, "username")
	valid.Required(password, "password")
	switch {
	case valid.HasErrors():
		fmt.Println("出错了！  " + valid.Errors[0].Key + valid.Errors[0].Message)
		c.TplName = "bad.html"
		return
	}
	u := &class.User{
		User_name: username,
		Password:  password,
	}
	err := u.Create()
	if err != nil {
		fmt.Println("这是error  ", err)
		c.TplName = "bad.html"
		return
	}
	c.TplName = "welcome.html"

}

func (c *UserController) Reallogin() {
	username := c.GetString("username")
	password := c.GetString("password")
	u := &class.User{
		User_name: username,
		Password:  password,
	}
	err := u.ReadDB()
	if err != nil {
		fmt.Println(err)
		c.TplName = "bad.html"
		return
	}
	// c.TplName = "welcome.html"
	mystruct := UserJson{
		code:    0,
		message: "成功",
		data: {
			"name": DataJson{},
		}
	}
	c.Data["json"] = &mystruct
	c.ServeJSON()
}
