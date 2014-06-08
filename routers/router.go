package routers

import (
	"github.com/astaxie/beego"
	"github.com/deffer/jimmyr/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
