package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Title"] = "Jobs"
	c.Layout = "layout.html"
	c.TplName = "index.tpl"
}
