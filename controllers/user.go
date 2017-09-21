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
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type DataJson struct {
}

func (c *UserController) PageLogin() {
	c.TplName = "login.html"
}

func (c *UserController) Register() {
	username := c.GetString("username")
	password := c.GetString("password")

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
	ms := make(map[string]interface {
	}, 0)
	ms["ss"] = 234234
	ms["map"] = map[string]interface{}{
		"name": 23,
		"time": "asf",
	}
	mystruct := &UserJson{
		Code:    1,
		Message: "成功",
		Data:    ms,
	}

	c.Data["json"] = mystruct
	c.ServeJSON()
}
