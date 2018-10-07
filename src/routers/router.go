package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //beego.Router("/index", &controllers.RootController{})
}
