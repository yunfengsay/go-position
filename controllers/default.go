package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *MainController) Post() {
	println("get post data")
	c.Data["Website"] = "POST"
	c.Data["Email"] = "post@gmail.com"
}
