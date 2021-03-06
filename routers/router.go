package routers

import (
	"github.com/astaxie/beego"
	"demo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/login", &controllers.UserController{})
	beego.Router("/user/register", &controllers.UserController{})
}
