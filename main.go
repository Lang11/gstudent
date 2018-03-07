package main

import (
	_ "github.com/Lang11/gstudent/routers"
	_ "github.com/Lang11/gstudent/db"
	"github.com/astaxie/beego"
	"github.com/Lang11/gstudent/controllers"
)

func main() {
	beego.SetLogFuncCall(true)
	beego.SetLevel(beego.LevelDebug)
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

