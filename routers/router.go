package routers

import (
	"github.com/Lang11/gstudent/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/create",&controllers.SController{},"post:Create")
	beego.Router("/update",&controllers.SController{},"post:Update")
	beego.Router("/read",&controllers.SController{},"post:Read")
	beego.Router("/delete",&controllers.SController{},"post:Delete")
	beego.Router("/upload",&controllers.SController{},"post:Upload")
	beego.Router("/jump",&controllers.SController{})



}
