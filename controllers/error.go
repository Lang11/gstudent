package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error500() {
	c.Data["content"] = "server error"
	c.TplName = "500.html"
}