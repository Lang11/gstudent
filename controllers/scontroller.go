package controllers

import "github.com/astaxie/beego"
import (
	"github.com/Lang11/gstudent/models"
	"html/template"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego/orm"
)

type SController struct {
	beego.Controller
}

func (c *SController) Create() {
	stu := new(models.GStudent)
	c.ParseForm(stu)
	beego.Info("Create:", stu)

	valid := validation.Validation{}
	b, err := valid.Valid(stu)
	if err != nil {
		beego.Error(err)

		c.Abort("500")
	}
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			beego.Error(err.Key, err.Message)
		}
		c.Abort("500")
	}

	o := orm.NewOrm()
	id, err := o.Insert(stu)
	if err == nil {
		beego.Info("S:", id)
		c.TplName = "200.html"
	}

}

func (c *SController) Update() {
	stu := new(models.GStudent)
	c.ParseForm(stu)
	beego.Info("Update:", stu)
}

func (c *SController) Delete() {
	stu := new(models.GStudent)
	c.ParseForm(stu)
	beego.Info("Delete:", stu)

	o := orm.NewOrm()
	if num,err:=o.Delete(&models.GStudent{Id:1});err== nil {
		beego.Info(num)
		c.TplName = "200.html"
	}else{
		c.Abort("500")
	}
}

func (c *SController) Read() {
	stu := new(models.GStudent)
	c.ParseForm(stu)
	beego.Info("Read:", stu)

	o := orm.NewOrm()
	id, e := c.GetInt("id")
	if e != nil {
		beego.Error(e)
		c.Abort("500")
	}

	s := models.GStudent{Id: id}
	err := o.Read(&s)

	if err == orm.ErrNoRows {
		beego.Warn("no row found")
	} else if err == orm.ErrMissPK {
		beego.Warn("no pk")
	} else {
		beego.Info(s.Name, s.Class)
	}

	c.TplName = "200.html"
}

func (c *SController) Upload() {
	f, h, err := c.GetFile("uploadname")
	if err != nil {
		beego.Error("getfile err ", err)
	}
	defer f.Close()
	c.SaveToFile("uploadname", "static/upload/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建

	c.TplName = "200.html"
}

func (c *SController) Get() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	switch url := c.GetString("url"); url {
	case "create":
		c.TplName = "create_action.html"
	case "update":
		c.TplName = "update_action.html"
	case "read":
		c.TplName = "read_action.html"
	case "delete":
		c.TplName = "delete_action.html"
	}

	//json := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	//c.Data["json"] = &json
	//c.ServeJSON()

	//xml :=`<?xml version="1.0" encoding="UTF-8"?><one><name>lk</name></one>`
	//c.Data["xml"] = &xml
	//c.ServeXML()

}
