package routers

import (
	"github.com/astaxie/beego"
	"demo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/login", &controllers.UserController{})
	beego.Router("/user/password", &controllers.UserController{})
	beego.Router("/user/register", &controllers.SignController{})
	beego.Router("/main/view", &controllers.MainViewController{})
	beego.Router("/examples/bootstrap_table/data/?:id",&controllers.DishController{})
}
